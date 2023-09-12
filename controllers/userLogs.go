package controllers

import (
	"gate/db"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	pocketModel "github.com/pocketbase/pocketbase/models"
)

func GetUserLogs(app *pocketbase.PocketBase) echo.HandlerFunc {
	return func(c echo.Context) error {

		// only admin endpoint 
		admin, _ := c.Get(apis.ContextAdminKey).(*pocketModel.Admin)
		if admin == nil {
			return apis.NewForbiddenError("access denied to this endpoint", nil)
		}

		userID := c.PathParam("user_id")
		page := c.PathParam("page")
		perPage := c.PathParam("per_page")

		pageInt, pageErr := strconv.Atoi(page)
		perPageInt, perPageErr := strconv.Atoi(perPage)

		if pageErr != nil || perPageErr != nil {
			return apis.NewBadRequestError("Invalid page or per_page parameter", nil)
		}

		// check if user exists 
		userExsits := db.CheckIfUserExists(userID , app.Dao())
		if !userExsits {
			return apis.NewBadRequestError("user doesnt exist", nil)
		}

		// get data 
		result , err := db.GetUserLogs(userID , pageInt , perPageInt , app.Dao())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error while processing the list of enteries"})
		}
		return c.JSON(http.StatusOK, result)
	}
}

func GetListOfAutorizedPersons(app *pocketbase.PocketBase) echo.HandlerFunc {
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


		// get the list of autorized persons
		result , err := db.GetAutorizedList(pageInt ,  perPageInt , app.Dao())

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error while processing the list of autorized persons"}) 
		}

		return c.JSON(http.StatusOK, result)

	}
}