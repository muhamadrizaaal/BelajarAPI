package routes

import (
	"BelajarAPI/config"
	"BelajarAPI/controller/user"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitRoute(c *echo.Echo, ctl user.UserController) {
	c.POST("/users", ctl.Register())
	c.POST("/login", ctl.Login())
	c.POST("/users/:user_id/activity", ctl.AddActivity(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
	c.PUT("/users/:user_id/activity/:id", ctl.UpdateActivity(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
	c.GET("/users/:user_id/activity", ctl.GetAllDataById(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
}
