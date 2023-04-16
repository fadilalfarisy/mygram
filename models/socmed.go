package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Socmed struct {
	GormModel
	Name       string `gorm:"not null" json:"name" form:"name" valid:"required"`
	Socmed_Url string `gorm:"not null" json:"socmed_url" form:"socmed_url" valid:"required"`
	UserID     uint
	User       *User
}

func (s *Socmed) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(s)
	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (s *Socmed) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(s)
	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
