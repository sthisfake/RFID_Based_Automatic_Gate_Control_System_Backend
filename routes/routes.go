package routes

import "github.com/pocketbase/pocketbase"

func RegisterAllRoutes(app *pocketbase.PocketBase) {
	gateRoutes(app)
	userData(app)
	dashbordRoutes(app)
}