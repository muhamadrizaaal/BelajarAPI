package main

import (
	"BelajarAPI/config"
	tControll "BelajarAPI/controller/todo"
	uControll "BelajarAPI/controller/user"
	"BelajarAPI/model/todo"
	"BelajarAPI/model/user"
	"BelajarAPI/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitSQL(cfg)

	// m := user.UserModel{Connection: db}
	// c := user.UserController{Model: m}
	m := user.UserModel{Connection: db}     // bagian yang menghungkan coding kita ke database / bagian dimana kita ngoding untk ke DB
	c := uControll.UserController{Model: m} // bagian yang menghandle segala hal yang berurusan dengan HTTP / echo
	tm := todo.TodoModel{Connection: db}
	tc := tControll.TodoController{Model: tm}
	routes.InitRoute(e, c, tc)
	e.Logger.Fatal(e.Start(":8000"))
}
