package models

import "github.com/pocketbase/pocketbase/models"

var _ models.Model = (*AuthorizedPerson)(nil)

type AuthorizedPerson struct {
	models.BaseModel
	FullName string `db:"full_name" json:"full_name"`
	TagId string `db:"tag_id" json:"tag_id"`
}

type EachHistory struct {
	Date string `json:"date"`
	EnterTime string `json:"enter_time"`
	ExistTime string `json:"exist_time"`
}

type UserHistory struct {
	Page int `json:"page"`
	PerPage int `json:"per_page"`
	TotalPages int `json:"total_pages"`
	TotalItems int `json:"total_items"`
	Items []EachHistory `json:"items"`
}

type AutorizedPersonsItem struct {
	UserId string `json:"user_id"`
	FullName string `json:"full_name"`
	LastEnteryDate string `json:"last_entery_date"`
	Status string `json:"status"`
}

type AutorizedPersonsList struct {
	Page int `json:"page"`
	PerPage int `json:"per_page"`
	TotalPages int `json:"total_pages"`
	TotalItems int `json:"total_items"`
	Items []AutorizedPersonsItem `json:"items"`
}

type PeopleInTheBuildingItem struct { 
	UserId string `json:"user_id"`
	FullName string `json:"full_name"`
	EnteryDate	string `json:"entery_date"`
	EnteryTime string `json:"entery_time"`
}

type PeopleInTheBuildingList struct {
	Page int `json:"page"`
	PerPage int `json:"per_page"`
	TotalPages int `json:"total_pages"`
	TotalItems int `json:"total_items"`
	Items []PeopleInTheBuildingItem `json:"items"`
}

func (m *AuthorizedPerson) TableName() string {
	return "authorized_person"
}