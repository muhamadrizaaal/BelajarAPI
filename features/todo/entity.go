package todo

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type TodoController interface {
	AddActivity() echo.HandlerFunc
	Update() echo.HandlerFunc
	GetData() echo.HandlerFunc
}

type TodoService interface {
	AddActivity(token *jwt.Token, kegiatan Todo) (Todo, error)
	UpdateActivity(newData Todo, idTodo uint, token *jwt.Token) error
	GetAllDataById(token *jwt.Token) ([]Todo, error)
}

type TodoModel interface {
	InsertActivity(hp string, kegiatan Todo) (Todo, error)
	UpdateActivity(newData Todo, idTodo uint, hp string) error
	GetAllDataById(hp string) ([]Todo, error)
}

type Todo struct {
	ID       uint
	Kegiatan string
}
