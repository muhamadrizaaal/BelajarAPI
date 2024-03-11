package user

import "github.com/labstack/echo/v4"

type UserController interface {
	SignUp() echo.HandlerFunc
	SignIn() echo.HandlerFunc
	// Add() echo.HandlerFunc
}

type UserService interface {
	Register(newData User) error
	Login(loginData User) (User, string, error)
}

type UserModel interface {
	InsertUser(newData User) error
	Login(hp string) (User, error)
}

type User struct {
	Nama     string
	Hp       string
	Password string
}

type Login struct {
	Hp       string `validate:"required,min=10,max=13,numeric"`
	Password string `validate:"required,alphanum,min=8"`
}

type Register struct {
	Nama     string `validate:"required,alpha"`
	Hp       string `validate:"required,min=10,max=13,numeric"`
	Password string `validate:"required,alphanum,min=8"`
}
