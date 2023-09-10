package controllers

import (
	"gate/db"
	"gate/utils"
	"net/http"
	"regexp"
	"strconv"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	pocketModel "github.com/pocketbase/pocketbase/models"
)

func GetDashbordView(app *pocketbase.PocketBase) echo.HandlerFunc {
	return func(c echo.Context) error {
		
		// only admin endpoint 
		admin, _ := c.Get(apis.ContextAdminKey).(*pocketModel.Admin)
		if admin == nil {
			return apis.NewForbiddenError("access denied to this endpoint", nil)
		}

		date := c.PathParam("date")
		page := c.PathParam("page")
		perPage := c.PathParam("per_page")

		pageInt, pageErr := strconv.Atoi(page)
		perPageInt, perPageErr := strconv.Atoi(perPage)

		if pageErr != nil || perPageErr != nil {
			return apis.NewBadRequestError("Invalid page or per_page parameter", nil)
		}

		pattern := `^\d{4}-\d{2}-\d{2}$`

		regex := regexp.MustCompile(pattern)
		
		if !regex.MatchString(date) {
			return apis.NewBadRequestError("Invalid date parameter", nil)
		}

		todayPersianDate , err  := utils.GeneratePersianDate(date)
		if err != nil {
			return apis.NewBadRequestError("Invalid date parameter2", nil)
		}
		utcMidnight := utils.PersianToUtc(todayPersianDate)

		// get the list result based on the date for dashbord
		result , err := db.GetDashbordResult(pageInt ,  perPageInt , app.Dao() , utcMidnight)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error while processing the list of autorized persons"}) 
		}

		return c.JSON(http.StatusOK, result)

	}
}