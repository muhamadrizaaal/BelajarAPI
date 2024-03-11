package data

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Kegiatan string
	Hp       string `gorm:"type:varchar(13);"`
}
