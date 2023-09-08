package main

import (
	"log"

	"github.com/pocketbase/pocketbase"
)

func main() {

	app := pocketbase.New()

	routes.RegisterAllRoutes(app)

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}