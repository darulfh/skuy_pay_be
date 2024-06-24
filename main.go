package main

import (
	"github.com/darulfh/skuy_pay_be/config"
	"github.com/darulfh/skuy_pay_be/database"
	"github.com/darulfh/skuy_pay_be/routes"
	m "github.com/darulfh/skuy_pay_be/usecase/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))

	routes.Routes(e, db)
	m.LogMiddlewares(e)

	// ====== HTTPS ========
	// httpsServer := &http.Server{
	// 	Addr:      fmt.Sprintf(":%s", config.AppConfig.AppPort),
	// 	Handler:   e,
	// 	TLSConfig: &tls.Config{},
	// }

	// if err := httpsServer.ListenAndServeTLS("server.crt", "server.key"); err != http.ErrServerClosed {
	// 	log.Fatal(err)
	// }

	// ====== HTTP ========
	e.Logger.Fatal(e.Start(":" + config.AppConfig.AppPort))
}
