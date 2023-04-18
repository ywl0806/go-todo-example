package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title       string    `gorm:"not null" json:"title"`
	Description string    `gorm:"not null" json:"descriptioin"`
	IsActive    bool      `gorm:"not null" json:"isActive"`
	Deadline    time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP" json:"deadLine"`
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

func UpdateTodo(db *gorm.DB, id string, Todo *Todo) (err error) {
	db.Where("ID = ?", id).First(&Todo)
	return nil
}
