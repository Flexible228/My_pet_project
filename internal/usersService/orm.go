package usersService

import (
	"My_pet_project/internal/web/tasks"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string       `json:"email"`
	Password string       `json:"password"`
	Id       int64        `json:"id"`
	Tasks    []tasks.Task `json:"tasks"`
}
