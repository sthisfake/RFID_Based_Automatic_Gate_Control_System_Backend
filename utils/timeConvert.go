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

func GetHistoryDate(t time.Time) string {
	irstTime := t.Add(time.Hour * 3).Add(time.Minute * 30) 
	pt := ptime.Unix(irstTime.Unix(), 0)
	return pt.Format("yyyy/MM/dd")
}

func GetHistoryTime(t time.Time) string {
	if  t.IsZero() {
		return ""
	}
	irstTime := t.Add(time.Hour * 3).Add(time.Minute * 30) 
	pt := ptime.Unix(irstTime.UTC().Unix(), 0)
	return pt.Format("hh:mm")
}

func GeneratePersianDate(date string ) (ptime.Time ,  error) {
	year , err := strconv.Atoi(date[0:4])
	if err != nil {
		return ptime.Time{}, err
	}
	month , err := strconv.Atoi(date[5:7])
	if err != nil {
		return ptime.Time{}, err
	}
	day , err := strconv.Atoi(date[8:10])
	if err != nil {
		return ptime.Time{}, err
	}
	pt := ptime.Date(year, ptime.Month(month) , day, 0, 0, 0, 0, ptime.Iran())
	return pt , nil
}

func PersianToUtc(timo ptime.Time) (time.Time) {
	t := timo.Time()
	newTime := t.Add(time.Hour * -24)
	theNewTime := time.Date(newTime.Year(), time.Month(newTime.Month()), newTime.Day() , 20 , 30 , 0 , 0 ,time.UTC )
	return theNewTime 
}