package main

import (
	"github.com/labstack/echo/v4"
	"github.com/rnwxyz/wishlist-sewa/config"
	"github.com/rnwxyz/wishlist-sewa/database"
	"github.com/rnwxyz/wishlist-sewa/routes"
)

func main() {
	config.InitConfig()

	db, err := database.ConnectDB()
	if err != nil {
		panic(err)
	}
	if err := database.MigrateDB(db); err != nil {
		panic(err)
	}

	e := echo.New()

	routes.InitRoutes(e, db)
	e.Logger.Fatal(e.Start(":" + config.Env.API_PORT))
}
