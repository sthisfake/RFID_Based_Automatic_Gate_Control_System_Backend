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

func TodayCarEnteryList(dao *daos.Dao , date time.Time) ([]*models.MonitoringTest, error) {
	
	entery := []*models.MonitoringTest{}

	err := dao.DB().NewQuery("SELECT * FROM monitoring WHERE created >= {:today} ").Bind(dbx.Params{
		"today" : date,
	}).All(&entery)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return entery, nil
}

func TodayCarExitList(dao *daos.Dao , date time.Time)  ([]*models.MonitoringTest, error) {
	
	exit := []*models.MonitoringTest{}

	err := dao.DB().NewQuery("SELECT * FROM monitoring WHERE created >= {:today} AND exit_time != {:null}").Bind(dbx.Params{
		"today" : date,
		"null" : "",
	}).All(&exit)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return exit, nil
}

func PeopleInTheBuilding(page int , per_page int , dao *daos.Dao) (models.PeopleInTheBuildingList, error) {
	var pageUsed int
	var perPageUsed int
	dbLogs := []*models.MonitoringTest{}
	result := models.PeopleInTheBuildingList{}

	err := dao.DB().NewQuery("SELECT * FROM monitoring WHERE exit_time = {:null}").Bind(dbx.Params{
		"null" : "",
	}).All(&dbLogs)

	if err != nil {
		fmt.Println(err.Error())
		return  models.PeopleInTheBuildingList{}, err
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
		result.Items = []models.PeopleInTheBuildingItem{}
		return result , nil
	}
	if (pageUsed == totalPage){
		something = dbLogs[(perPageUsed * (totalPage - 1)) :  ]
	}
	if (pageUsed  < totalPage){
		something = dbLogs[ (pageUsed -1 ) * perPageUsed :  pageUsed * perPageUsed ]
	}

	for _ , item := range(something) {

		fullName , err  := GetNameWithUser(dao , item.UserId)
		if err != nil {
			return  models.PeopleInTheBuildingList{}, err
		}

		eachItem := models.PeopleInTheBuildingItem{
			UserId: item.UserId ,
			FullName: fullName,
			EnteryDate: utils.GetHistoryDate(item.EnteryTime.Time()),
			EnteryTime: utils.GetHistoryTime(item.EnteryTime.Time()),
		}
		result.Items = append(result.Items, eachItem)
	} 

	return result , nil

}

func MonitorQuerry(dao *daos.Dao) *dbx.SelectQuery {
	return dao.ModelQuery(&models.MonitoringTest{})
}

func GetNameWithUser(dao *daos.Dao , userId string )  (string, error){
	
	dbLogs := models.AuthorizedPerson{}
	err := UserQuerry(dao).AndWhere(dbx.HashExp{"id": userId}).One(&dbLogs)
	if err != nil {
		fmt.Println(err.Error())
		return  "", err
	}

	return dbLogs.FullName , nil
}
