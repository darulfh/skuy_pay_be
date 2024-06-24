package main

import (
	"github.com/darulfh/skuy_pay_be/config"
	"github.com/darulfh/skuy_pay_be/database"
	"github.com/darulfh/skuy_pay_be/routes"
	m "github.com/darulfh/skuy_pay_be/usecase/middlewares"

	"github.com/labstack/echo/v4"
)

func main() {

	config.LoadConfig()

	db, err := database.ConnectDB()
	if err != nil {
		panic(err)
	}

	// database.Drop(db)

	database.Migrate(db)

	e := echo.New()
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"*"},
	// 	AllowHeaders: []string{"*"},
	// }))

	routes.Routes(e, db)
	m.LogMiddlewares(e)

	// ====== HTTP ========
	e.Logger.Fatal(e.Start(":" + config.AppConfig.AppPort))
}
