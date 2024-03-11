package data

import (
	"BelajarAPI/features/user"
	"errors"
	"log"

	"gorm.io/gorm"
)

type model struct {
	connection *gorm.DB
}

func New(db *gorm.DB) user.UserModel {
	return &model{
		connection: db,
	}
}

func (m *model) InsertUser(newData user.User) error {
	err := m.connection.Create(&newData).Error

	if err != nil {
		defer func() {
			if err := recover(); err != nil {
				log.Println("error database process:", err)

			}
		}()
		return errors.New("terjadi masalah pada database")
	}

	return nil
}

func (m *model) Login(hp string) (user.User, error) {
	var result user.User
	if err := m.connection.Model(&User{}).Where("hp = ? ", hp).First(&result).Error; err != nil {
		return user.User{}, err
	}
	return result, nil
}
