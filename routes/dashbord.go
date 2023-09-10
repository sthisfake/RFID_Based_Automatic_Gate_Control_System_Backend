package routes

import (
	"gate/controllers"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func dashbordRoutes(app *pocketbase.PocketBase) {
	// get the full list of cities
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.AddRoute(echo.Route{
			Method:  http.MethodGet,
			Path:    "/dashbord/:date/:page/:per_page/",
			Handler: controllers.GetDashbordView(app),
		})
		return nil
	})

}