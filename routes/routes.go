package routes

import (
	"BelajarAPI/config"
	todo "BelajarAPI/features/todo"
	user "BelajarAPI/features/user"

	// echojwt "github.com/labstack/echo-jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitRoute(c *echo.Echo, ctl user.UserController, tc todo.TodoController) {
	userRoute(c, ctl)
	todoRoute(c, tc)

}

func userRoute(c *echo.Echo, ctl user.UserController) {
	c.POST("/users", ctl.SignUp())
	c.POST("/login", ctl.SignIn())
}

func todoRoute(c *echo.Echo, tc todo.TodoController) {
	c.POST("/activity", tc.AddActivity(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
	c.PUT("/activity/:id", tc.Update(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
	c.GET("/activity", tc.GetData(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
}
