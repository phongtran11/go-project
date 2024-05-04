package main

import (
	"github.com/phongtran11/go-project/database"
	"github.com/phongtran11/go-project/routes"
)

func main() {
	// init connection to database
	database.ConnectDB()

	// init routes
	routes.InitRoutes()
}
