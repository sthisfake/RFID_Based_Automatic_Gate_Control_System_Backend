package utils

import (
	"strconv"
	"time"

	ptime "github.com/yaa110/go-persian-calendar"
)

func ConvertTime(t time.Time) (time.Time , time.Time) {
	irstTime := t.Add(time.Hour * 3).Add(time.Minute * 30) 
	newTime := irstTime.Add(time.Hour * -24)
	theNewTime := time.Date(newTime.Year(), time.Month(newTime.Month()), newTime.Day() , 20 , 30 , 0 , 0 ,time.UTC )
	return theNewTime , irstTime
}


func GetFullDateToday(t time.Time) string {
	pt := ptime.New(t)
	day := pt.Weekday().String() + " " +  strconv.Itoa(pt.Day()) + " " +  pt.Month().String() + " " +  strconv.Itoa(pt.Year())
	return day
}