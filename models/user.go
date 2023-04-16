package models

import (
	"challenge-4/helpers"
	"errors"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// User represents the model for an user
type User struct {
	GormModel
	Email    string    `gorm:"not null; uniqueIndex:unique_index" json:"email" form:"email" valid:"required, email"`
	Username string    `gorm:"not null; uniqueIndex:unique_index" json:"username" form:"username" valid:"required"`
	Password string    `gorm:"not null" json:"password" form:"password" valid:"required, minstringlength(6)"`
	Age      int       `gorm:"not null" json:"age" form:"age" valid:"required"`
	Photo    []Photo   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"photos"`
	Comment  []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comments"`
	Socmed   []Socmed  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"socmed"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)
	if errCreate != nil {
		err = errCreate
		return
	}

	if u.Age < 8 {
		err = errors.New("age must greater than equal 8")
		return
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}
