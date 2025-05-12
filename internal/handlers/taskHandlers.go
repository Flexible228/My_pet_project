package handlers

import (
	"My_pet_project/internal/tasksService"
	"My_pet_project/internal/web/tasks"
	"golang.org/x/net/context"
)

type TasksHandler struct {
	Service *tasksService.TaskService
}

func NewTasksHandler(service *tasksService.TaskService) *TasksHandler {
	return &TasksHandler{
		Service: service,
	}
}

func (h *TasksHandler) GetTasksUserUserId(_ context.Context, req tasks.GetTasksUserUserIdRequestObject) (tasks.GetTasksUserUserIdResponseObject, error) {
	allTasks, err := h.Service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	response := tasks.GetTasksUserUserId200JSONResponse{}

	for _, tsk := range allTasks {
		if tsk.UserId == req.UserId {
			task := tasks.Task{
				Id:     &tsk.Id,
				Task:   &tsk.Task,
				IsDone: &tsk.IsDone,
			}
			response = append(response, task)
		}
	}

	return response, nil
}

func (h *TasksHandler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.Service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	response := tasks.GetTasks200JSONResponse{}

	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.Id,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
		}
		response = append(response, task)
	}

	return response, nil
}

func (h *TasksHandler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	// Распаковываем тело запроса напрямую, без декодера!
	taskRequest := request.Body
	// Обращаемся к сервису и создаем задачу
	taskToCreate := tasksService.Task{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
		UserId: *taskRequest.UserId,
	}
	createdTask, err := h.Service.CreateTask(taskToCreate, *taskRequest.UserId)

	if err != nil {
		return nil, err
	}
	// создаем структуру респонс
	response := tasks.PostTasks201JSONResponse{
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
		Id:     &createdTask.Id,
	}
	// Просто возвращаем респонс!
	return response, nil
}

func (h *TasksHandler) DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	// Удаляем задачу по ID
	err := h.Service.DeleteTaskByID(uint(request.Id))
	if err != nil {
		// Обработайте ошибку - например, если задача не найдена
		return nil, err
	}

	// Возвращаем успешный ответ (можно вернуть пустой ответ)
	return tasks.DeleteTasksId204Response{}, nil
}

func (h *TasksHandler) PatchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	// Получение данных для обновления
	updatedTask := tasksService.Task{}
	if request.Body.Task != nil {
		updatedTask.Task = *request.Body.Task
	}
	if request.Body.IsDone != nil {
		updatedTask.IsDone = *request.Body.IsDone
	}

	// Обновляем задачу
	taskId := request.Id
	updatedTask, err := h.Service.UpdateTaskByID(uint(taskId), updatedTask)
	if err != nil {
		// Обработайте ошибку - например, если задача не найдена
		return nil, err
	}

	// Успешно обновленная задача
	response := tasks.PatchTasksId200JSONResponse{
		Id:     &updatedTask.Id,
		Task:   &updatedTask.Task,
		IsDone: &updatedTask.IsDone,
	}
	return response, nil
}
