package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string	
	Password string
	Super    bool
}