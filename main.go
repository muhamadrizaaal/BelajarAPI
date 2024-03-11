package main

import (
	"BelajarAPI/config"
	td "BelajarAPI/features/todo/data"
	th "BelajarAPI/features/todo/handler"
	ts "BelajarAPI/features/todo/services"
	"BelajarAPI/features/user/data"
	"BelajarAPI/features/user/handler"
	"BelajarAPI/features/user/services"

	"BelajarAPI/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()            // inisiasi echo
	cfg := config.InitConfig() // baca seluruh system variable
	db := config.InitSQL(cfg)  // konek DB

	userData := data.New(db)
	userService := services.NewService(userData)
	userHandler := handler.NewUserHandler(userService)

	todoData := td.New(db)
	todoService := ts.NewTodoService(todoData)
	todoHandler := th.NewHandler(todoService)

	// e.Pre(middleware.RemoveTrailingSlash())
	// e.Use(middleware.Logger())
	// e.Use(middleware.CORS()) // ini aja cukup
	routes.InitRoute(e, userHandler, todoHandler)
	e.Logger.Fatal(e.Start(":8000"))
}
