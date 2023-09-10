package db

import (
	"gate/models"
	"gate/utils"
	"math"
	"time"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
)

func GetDashbordResult(page int, per_page int, dao *daos.Dao , date time.Time) (models.DashbordMonitoringList, error) {

	var pageUsed int
	var perPageUsed int
	dbLogs := []*models.MonitoringTest{}
	result := models.DashbordMonitoringList{}

	err := dao.DB().NewQuery("SELECT * FROM monitoring WHERE created >= {:today} ").Bind(dbx.Params{
		"today" : date,
	}).All(&dbLogs)

	if err != nil {
		return models.DashbordMonitoringList{}, err
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
		result.Items = []models.DashbordMonitoringItem{}
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
			return  models.DashbordMonitoringList{}, err
		}

		eachItem := models.DashbordMonitoringItem{
			UserId: item.UserId ,
			FullName: fullName,
			ExitTime : utils.GetHistoryTime(item.ExitTime.Time()),
			EnteryTime: utils.GetHistoryTime(item.EnteryTime.Time()),
		}
		result.Items = append(result.Items, eachItem)
	} 

	return result , nil
}