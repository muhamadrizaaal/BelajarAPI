package model

import (
	"gorm.io/gorm"
)

type User struct {
	UserID   uint   `json:"userid" form:"userid" gorm:"primaryKey"`
	Hp       string `json:"hp" form:"hp" validate:"required,max=13,min=10"`
	Nama     string `json:"nama" form:"nama" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type DaftarKegiatan struct {
	gorm.Model
	Kegiatan string
	UserID   uint
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

func (um *UserModel) GetLastUserID() (uint, error) {
	var lastUser User

	// Query untuk mendapatkan pengguna terakhir berdasarkan user_id terbesar
	if err := um.Connection.Order("user_id desc").First(&lastUser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// Tabel kosong, return 0 sebagai UserID pertama
			return 0, nil
		}
		return 0, err
	}

	return lastUser.UserID, nil
}

func (um *UserModel) AddActivity(newData DaftarKegiatan, userid uint64) error {

	if err := um.Connection.Table("daftar_kegiatans").Where("user_id = ?", userid).Create(&newData).Error; err != nil {
		return err
	}

	// err := um.Connection.Create(&newData).Error
	// if err != nil {
	// 	return err
	// }
	return nil
}

func (um *UserModel) UpdateActivity(newData DaftarKegiatan, userid uint, id uint) error {
	if err := um.Connection.Model(DaftarKegiatan{}).Where("user_id = ? AND id = ?", userid, id).Update("kegiatan", newData.Kegiatan).Error; err != nil {
		return err
	}
	return nil
}

func (um *UserModel) GetAllDataById(userid uint64) ([]DaftarKegiatan, error) {
	var result []DaftarKegiatan
	if err := um.Connection.Where("user_id = ?", userid).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
