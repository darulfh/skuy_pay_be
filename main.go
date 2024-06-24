package main

import (
	"net/http"

	"github.com/darulfh/skuy_pay_be/config"
	"github.com/darulfh/skuy_pay_be/model"
	echo "github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, model.MetaData{
			Message: "Success",
		})
	})

	e.Logger.Fatal(e.Start(":" + config.AppConfig.AppPort))

	// config.LoadConfig()

	// db, err := database.ConnectDB()
	// if err != nil {
	// 	panic(err)
	// }

	// // database.Drop(db)

	// database.Migrate(db)

	// e := echo.New()
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"*"},
	// 	AllowHeaders: []string{"*"},
	// }))

	// routes.Routes(e, db)
	// m.LogMiddlewares(e)

	// // ====== HTTP ========
	// e.Logger.Fatal(e.Start(":" + config.AppConfig.AppPort))
}
