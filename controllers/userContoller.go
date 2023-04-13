package controllers

import "gorm.io/gorm"

type UserRepo struct {
	Db *gorm.DB
}
