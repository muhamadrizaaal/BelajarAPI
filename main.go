package main

import (
	"BelajarAPI/config"
	"BelajarAPI/controller/user"
	"BelajarAPI/model"
	"BelajarAPI/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitSQL(cfg)

	m := model.UserModel{Connection: db}
	c := user.UserController{Model: m}
	routes.InitRoute(e, c)
	e.Logger.Fatal(e.Start(":8000"))
}
