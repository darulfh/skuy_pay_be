package main

import (
	"BE-Golang/config"
	"BE-Golang/database"
	"BE-Golang/routes"
	m "BE-Golang/usecase/middlewares"

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
