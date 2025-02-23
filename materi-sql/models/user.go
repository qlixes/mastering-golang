package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `json:"name"`
	Email string `json:"email" gorm:"unique"`
	Password string `json:"password"`	
	Roles []Role `json:"roles" gorm:"many2many:user_roles"`
}

type Role struct {
	gorm.Model
	Name string `json:"name"`
	Users []User `json:"users" gorm:"many2many:user_roles"`
}