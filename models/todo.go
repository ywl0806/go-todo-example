package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Discription string
	IsActive    bool `gorm:"not null"`
	Deadline    time.Time
}

func CreateTodo(db *gorm.DB, Todo *Todo) (err error) {
	err = db.Create(Todo).Error
	if err != nil {
		return err
	}
	return nil
}

func GetTodos(db *gorm.DB, Todo *[]Todo) (err error) {
	err = db.Find(Todo).Error
	if err != nil {
		return err
	}
	return nil
}
