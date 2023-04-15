package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title       string    `gorm:"not null"`
	Discription string    `gorm:"not null"`
	IsActive    bool      `gorm:"not null"`
	Deadline    time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
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

func UpdateTodos(db *gorm.DB, Todo *Todo) (err error) {
	return nil
}
