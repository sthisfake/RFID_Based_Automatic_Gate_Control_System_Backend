package routes

import (
	"gate/controllers"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func gateRoutes(app *pocketbase.PocketBase) {
	// get the full list of cities
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.AddRoute(echo.Route{
			Method:  http.MethodGet,
			Path:    "/overall_states/today",
			Handler: controllers.GetOverallViewToday(app),
		})
		return nil
	})

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.AddRoute(echo.Route{
			Method:  http.MethodGet,
			Path:    "/people_in_building/:page/:per_page/",
			Handler: controllers.GetCurrentPeopleInBuilding(app),
		})
		return nil
	})
}