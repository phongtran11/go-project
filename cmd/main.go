package main

import (
	"github.com/joho/godotenv"
	"github.com/phongtran11/go-project/database"
	"github.com/phongtran11/go-project/routes"
)

func main() {
	// load dotenv
	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}

	// init connection to database
	database.ConnectDB()

	// init routes
	routes.InitRoutes()
}
