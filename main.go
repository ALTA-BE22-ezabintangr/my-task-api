package main

import (
	"myTaskApp/app/configs"
	"myTaskApp/app/databases"
	"myTaskApp/app/migrations"
	"myTaskApp/app/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := configs.InitConfig()
	dbMySql := databases.InitDBMysql(cfg)
	// dbMysql := databases.InitDBPosgres(cfg)

	// create new instance echo
	e := echo.New()

	migrations.InitialMigration()
	routes.InitRouter(e, dbMySql)

	// start server and port
	e.Logger.Fatal(e.Start(":8080"))
}
