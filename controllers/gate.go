package controllers

import (
	"gate/db"
	"gate/models"
	"gate/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	pocketModel "github.com/pocketbase/pocketbase/models"
)

func GetOverallViewToday(app *pocketbase.PocketBase) echo.HandlerFunc {
	return func(c echo.Context) error {

		// only admin endpoint 
		admin, _ := c.Get(apis.ContextAdminKey).(*pocketModel.Admin)
		if admin == nil {
			return apis.NewForbiddenError("access denied to this endpoint", nil)
		}

		currentTime := time.Now().UTC()
		midNightTime , iranDate := utils.ConvertTime(currentTime)

		
		enteriesList ,  err := db.TodayCarEnteryList(app.Dao() ,midNightTime )
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error while processing the list of enteries"})
		}

		exitList ,  err := db.TodayCarExitList(app.Dao() ,midNightTime )
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error while processing the list of enteries"})
		}

		response := models.MonitoringOverall{}
		response.Entery = strconv.Itoa(len(enteriesList))
		response.Exit = strconv.Itoa(len(exitList))
		response.Day = utils.GetFullDateToday(iranDate)

		return c.JSON(http.StatusOK, response)

	}
}


func GetCurrentPeopleInBuilding(app *pocketbase.PocketBase)echo.HandlerFunc{
	return func(c echo.Context) error {

		// only admin endpoint 
		admin, _ := c.Get(apis.ContextAdminKey).(*pocketModel.Admin)
		if admin == nil {
			return apis.NewForbiddenError("access denied to this endpoint", nil)
		}

		page := c.PathParam("page")
		perPage := c.PathParam("per_page")

		pageInt, pageErr := strconv.Atoi(page)
		perPageInt, perPageErr := strconv.Atoi(perPage)

		if pageErr != nil || perPageErr != nil {
			return apis.NewBadRequestError("Invalid page or per_page parameter", nil)
		}

		// get the list of people in the building
		result , err := db.PeopleInTheBuilding(pageInt ,  perPageInt , app.Dao())

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error while processing the list of autorized persons"}) 
		}

		return c.JSON(http.StatusOK, result)
	}
}