package data

import (
	"BelajarAPI/features/todo"
	"errors"

	"gorm.io/gorm"
)

type model struct {
	connection *gorm.DB
}

func New(db *gorm.DB) todo.TodoModel {
	return &model{
		connection: db,
	}
}

func (m *model) InsertActivity(hp string, kegiatan todo.Todo) (todo.Todo, error) {
	var inputProcess = Todo{Kegiatan: kegiatan.Kegiatan, Hp: hp}
	if err := m.connection.Create(&inputProcess).Error; err != nil {
		return todo.Todo{}, err
	}

	return todo.Todo{Kegiatan: kegiatan.Kegiatan}, nil
}

func (m *model) UpdateActivity(newData todo.Todo, idTodo uint, hp string) error {
	var query = m.connection.Where("id = ? AND hp = ?", idTodo, hp).Updates(newData)
	if err := query.Error; err != nil {
		return err
	}

	if query.RowsAffected < 1 {
		return errors.New("no data affected")
	}

	return nil
}

func (m *model) GetAllDataById(hp string) ([]todo.Todo, error) {
	var result []todo.Todo
	if err := m.connection.Where("hp = ?", hp).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
