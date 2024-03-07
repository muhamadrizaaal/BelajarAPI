package user

import (
	"BelajarAPI/model/todo"

	"gorm.io/gorm"
)

type User struct {
	Hp       string `gorm:"primaryKey"`
	Nama     string
	Password string
	Todos    []todo.Todo `gorm:"foreignKey:Hp;references:Hp"`
}

type UserModel struct {
	Connection *gorm.DB
}

func (um *UserModel) Register(newData User) error {
	err := um.Connection.Create(&newData).Error
	if err != nil {
		return err
	}
	return nil
}

func (um *UserModel) Login(hp string, password string) (User, error) {
	var result User
	if err := um.Connection.Where("hp = ? AND password = ?", hp, password).First(&result).Error; err != nil {
		return User{}, err
	}

	return result, nil
}

// func (um *UserModel) GetLastUserID() (uint, error) {
// 	var lastUser User

// 	// Query untuk mendapatkan pengguna terakhir berdasarkan user_id terbesar
// 	if err := um.Connection.Order("user_id desc").First(&lastUser).Error; err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			// Tabel kosong, return 0 sebagai UserID pertama
// 			return 0, nil
// 		}
// 		return 0, err
// 	}

// 	return lastUser.UserID, nil
// }
