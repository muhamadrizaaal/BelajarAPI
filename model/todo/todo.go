package todo

import (
	"errors"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Kegiatan string
	Hp       string
}

type TodoModel struct {
	Connection *gorm.DB
}

func (ut *TodoModel) AddActivity(newData Todo) error {

	if err := ut.Connection.Create(&newData).Error; err != nil {
		return err
	}
	return nil
}

func (ut *TodoModel) UpdateActivity(newData Todo, id uint, hp string) error {
	var query = ut.Connection.Where("id = ? AND hp = ?", id, hp).Updates(newData)
	if err := query.Error; err != nil {
		return err
	}

	if query.RowsAffected < 1 {
		return errors.New("no data affected")
	}

	return nil
}

func (ut *TodoModel) GetAllDataById(hp string) ([]Todo, error) {
	var result []Todo
	if err := ut.Connection.Where("hp = ?", hp).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
