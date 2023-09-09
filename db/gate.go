package db

import (
	"fmt"
	"gate/models"
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

func MonitorQuerry(dao *daos.Dao) *dbx.SelectQuery {
	return dao.ModelQuery(&models.MonitoringTest{})
}
