package routes

import (
	"BelajarAPI/config"
	"BelajarAPI/controller/todo"
	"BelajarAPI/controller/user"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitRoute(c *echo.Echo, ctl user.UserController, tc todo.TodoController) {
	userRoute(c, ctl)
	todoRoute(c, tc)

}

func userRoute(c *echo.Echo, ctl user.UserController) {
	c.POST("/users", ctl.Register())
	c.POST("/login", ctl.Login())
}

func todoRoute(c *echo.Echo, tc todo.TodoController) {
	c.POST("/activity", tc.AddActivity(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
	c.PUT("/activity/:id", tc.UpdateActivity(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
	c.GET("/activity", tc.GetAllDataById(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
}
