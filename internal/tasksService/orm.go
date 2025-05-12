package tasksService

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
	Id     int64  `json:"id"`
	UserId uint   `json:"user_id"`
}
