package db

import (
	"fmt"
	"gate/models"
	"gate/utils"
	"math"
	"time"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
)

func CheckIfUserExists(user_id string , dao *daos.Dao) (bool) {

	user := models.AuthorizedPerson{}

	err := UserQuerry(dao).AndWhere(dbx.HashExp{"id": user_id}).Limit(1).One(&user)

	return err == nil

}

func GetUserLogs(user_id string , page int , per_page int ,  dao *daos.Dao) (models.UserHistory, error) {

	var pageUsed int
	var perPageUsed int
	dbLogs := []*models.MonitoringTest{}
	result := models.UserHistory{}

	err := MonitorQuerry(dao).AndWhere(dbx.HashExp{"user_id": user_id}).All(&dbLogs)
	if err != nil {
		fmt.Println(err.Error())
		return  models.UserHistory{}, err
	}

	result.TotalItems = len(dbLogs)

	if page <= 0 {
		pageUsed = 1
    }else {
		pageUsed = page
	}

	if per_page <= 0 {
		perPageUsed = 10
	}else{
		perPageUsed = per_page
	}

	totalPage := int(math.Ceil(float64(len(dbLogs)) / float64( perPageUsed)))
	result.TotalPages = totalPage
	result.Page = pageUsed
	result.PerPage = perPageUsed
	something := []*models.MonitoringTest{} 

	if ( pageUsed  > totalPage){
		result.Items = []models.EachHistory{}
		return result , nil
	}
	if (pageUsed == totalPage){
		something = dbLogs[(perPageUsed * (totalPage - 1)) :  ]
	}
	if (pageUsed  < totalPage){
		something = dbLogs[ (pageUsed -1 ) * perPageUsed :  pageUsed * perPageUsed ]
	}

	for _ , item := range(something) {
		eachItem := models.EachHistory{
			Date: utils.GetHistoryDate(item.Created.Time()),
			EnterTime: utils.GetHistoryTime(item.EnteryTime.Time()),
			ExistTime: utils.GetHistoryTime(item.ExitTime.Time()),
		}
		result.Items = append(result.Items, eachItem)
	} 

	return result , nil
}


func GetAutorizedList(page int , per_page int , dao *daos.Dao) (models.AutorizedPersonsList, error) {
	var pageUsed int
	var perPageUsed int
	dbLogs := []*models.AuthorizedPerson{}
	result := models.AutorizedPersonsList{}

	err := UserQuerry(dao).All(&dbLogs)
	if err != nil {
		fmt.Println(err.Error())
		return  models.AutorizedPersonsList{}, err
	}

	result.TotalItems = len(dbLogs)

	if page <= 0 {
		pageUsed = 1
    }else {
		pageUsed = page
	}

	if per_page <= 0 {
		perPageUsed = 10
	}else{
		perPageUsed = per_page
	}

	totalPage := int(math.Ceil(float64(len(dbLogs)) / float64( perPageUsed)))
	result.TotalPages = totalPage
	result.Page = pageUsed
	result.PerPage = perPageUsed
	something := []*models.AuthorizedPerson{} 

	if ( pageUsed  > totalPage){
		result.Items = []models.AutorizedPersonsItem{}
		return result , nil
	}
	if (pageUsed == totalPage){
		something = dbLogs[(perPageUsed * (totalPage - 1)) :  ]
	}
	if (pageUsed  < totalPage){
		something = dbLogs[ (pageUsed -1 ) * perPageUsed :  pageUsed * perPageUsed ]
	}

	for _ , item := range(something) {
		lastEntery , err  := GetLastEnteryDate(dao , item.Id)
		if err != nil {
			return  models.AutorizedPersonsList{}, err
		}
		status , err := GetUserStatus(dao , item.Id)
		if err != nil {
			return  models.AutorizedPersonsList{}, err
		}
		eachItem := models.AutorizedPersonsItem{
			UserId: item.Id ,
			FullName: item.FullName,
			LastEnteryDate: lastEntery,
			Status: status,
		}
		result.Items = append(result.Items, eachItem)
	} 

	return result , nil

}

func GetLastEnteryDate(dao *daos.Dao , userId string) (string , error) {

	dbLogs := []*models.MonitoringTest{}
	err := MonitorQuerry(dao).AndWhere(dbx.HashExp{"user_id": userId}).AndOrderBy("created DESC").All(&dbLogs)
	if err != nil {
		fmt.Println(err.Error())
		return  "", err
	}

	if len(dbLogs) == 0 {
		return "" , nil
	}

	date := utils.GetHistoryDate(dbLogs[0].Created.Time())
	return date , nil

}

func GetUserStatus(dao *daos.Dao , userId string) (string , error) {

	dbLogs := []*models.MonitoringTest{}
	currentTime := time.Now().UTC()
	midNightTime , _ := utils.ConvertTime(currentTime)

	err := dao.DB().NewQuery("SELECT * FROM monitoring WHERE created >= {:today} AND user_id = {:user_id} ").Bind(dbx.Params{
		"today" : midNightTime,
		"user_id" : userId,
	}).All(&dbLogs)

	if err != nil {
		return "", err
	}

	if len(dbLogs) == 0 {
		return "غیر فعال" , nil
	}

	return "فعال" , nil
}


func UserQuerry(dao *daos.Dao) *dbx.SelectQuery {
	return dao.ModelQuery(&models.AuthorizedPerson{})
}
