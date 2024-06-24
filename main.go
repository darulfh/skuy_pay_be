package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})

	fmt.Println("Serving on port 8080")
	http.ListenAndServe(":8080", nil)

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
