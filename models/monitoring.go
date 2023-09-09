package models

import (
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tools/types"
)

var _ models.Model = (*MonitoringTest)(nil)

type MonitoringTest struct {
	models.BaseModel
	EnteryTime types.DateTime `db:"entery_time" json:"entery_time"`
}

type MonitoringOverall struct{
	Day string `json:"day"`
	Entery string `json:"entery"`
	Exit string `json:"exit"`
}

func (m *MonitoringTest) TableName() string {
	return "monitoring"
}