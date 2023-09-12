package routes

import (
	"gate/controllers"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func userData(app *pocketbase.PocketBase) {
	// get the full list of cities
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.AddRoute(echo.Route{
			Method:  http.MethodGet,
			Path:    "/user_logs/:user_id/:page/:per_page/",
			Handler: controllers.GetUserLogs(app),
		})
		return nil
	})

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.AddRoute(echo.Route{
			Method:  http.MethodGet,
			Path:    "/autorized_persons_list/:page/:per_page/",
			Handler: controllers.GetListOfAutorizedPersons(app),
		})
		return nil
	})
}