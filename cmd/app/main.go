package main

import (
	"My_pet_project/internal/database"
	"My_pet_project/internal/handlers"
	"My_pet_project/internal/tasksService"
	"My_pet_project/internal/usersService"
	"My_pet_project/internal/web/tasks"
	"My_pet_project/internal/web/users"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	database.InitDB()
	err := database.DB.AutoMigrate(&tasksService.Task{})
	if err != nil {
		return
	}

	tasksRepo := tasksService.NewTaskRepository(database.DB)
	TasksService := tasksService.NewTasksService(tasksRepo)
	tasksHandler := handlers.NewTasksHandler(TasksService)

	usersRepo := usersService.NewUsersRepository(database.DB)
	usersService := usersService.NewUsersService(usersRepo)
	usersHandler := handlers.NewUsersHandler(usersService)

	// Инициализируем echo
	e := echo.New()

	// используем Logger и Recover
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Прикол для работы в echo. Передаем и регистрируем хендлер в echo
	strictTasksHandler := tasks.NewStrictHandler(tasksHandler, nil) // тут будет ошибка
	tasks.RegisterHandlers(e, strictTasksHandler)

	strictUsersHandler := users.NewStrictHandler(usersHandler, nil)
	users.RegisterHandlers(e, strictUsersHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
