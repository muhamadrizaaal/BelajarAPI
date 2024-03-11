package services

import (
	"BelajarAPI/features/todo"
	"BelajarAPI/helper"
	"BelajarAPI/middleware"
	"errors"
	"log"

	// "net/http"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	// "gorm.io/gorm"
)

type service struct {
	m todo.TodoModel
	v *validator.Validate
}

func NewTodoService(model todo.TodoModel) todo.TodoService {
	return &service{
		m: model,
		v: validator.New(),
	}
}

func (s *service) AddActivity(token *jwt.Token, kegiatan todo.Todo) (todo.Todo, error) {
	hp := middleware.DecodeToken(token)
	if hp == "" {
		log.Println("error decode token:", "token tidak ditemukan")
		return todo.Todo{}, errors.New("data tidak valid")
	}

	err := s.v.Struct(&kegiatan)
	if err != nil {
		log.Println("error validasi", err.Error())
		return todo.Todo{}, err
	}

	result, err := s.m.InsertActivity(hp, kegiatan)
	if err != nil {
		return todo.Todo{}, errors.New(helper.ServerGeneralError)
	}

	return result, nil
}

func (s *service) UpdateActivity(newData todo.Todo, idTodo uint, token *jwt.Token) error {
	hp := middleware.DecodeToken(token)
	if hp == "" {
		log.Println("error decode token:", "token tidak ditemukan")
		return errors.New("data tidak valid")
	}

	err := s.v.Struct(&newData)
	if err != nil {
		log.Println("error validasi", err.Error())
		return err
	}

	error := s.m.UpdateActivity(newData, idTodo, hp)
	if error != nil {
		return errors.New(helper.ServerGeneralError)
	}

	return nil
}

func (s *service) GetAllDataById(token *jwt.Token) ([]todo.Todo, error) {
	hp := middleware.DecodeToken(token)
	if hp == "" {
		log.Println("error decode token:", "token tidak ditemukan")
		return nil, errors.New("data tidak valid")
	}

	result, err := s.m.GetAllDataById(hp)
	if err != nil {
		return nil, errors.New(helper.ServerGeneralError)
	}

	return result, nil
}
