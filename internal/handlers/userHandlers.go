package handlers

import (
	"My_pet_project/internal/usersService"
	"My_pet_project/internal/web/users"
	"golang.org/x/net/context"
)

type UsersHandler struct {
	Service *usersService.UserService
}

func NewUsersHandler(service *usersService.UserService) *UsersHandler {
	return &UsersHandler{
		Service: service,
	}
}

func (h *UsersHandler) GetUsers(_ context.Context, _ users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	// Получение всех задач из сервиса
	allUsers, err := h.Service.GetAllUsers()
	if err != nil {
		return nil, err
	}

	// Создаем переменную респон типа 200джейсонРеспонс
	// Которую мы потом передадим в качестве ответа
	response := users.GetUsers200JSONResponse{}

	// Заполняем слайс response всеми задачами из БД
	for _, usr := range allUsers {
		user := users.User{
			Id:       &usr.Id,
			Email:    &usr.Email,
			Password: &usr.Password,
		}
		response = append(response, user)
	}

	// САМОЕ ПРЕКРАСНОЕ. Возвращаем просто респонс и nil!
	return response, nil
}

func (h *UsersHandler) PostUsers(_ context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	// Распаковываем тело запроса напрямую, без декодера!
	userRequest := request.Body
	// Обращаемся к сервису и создаем задачу
	userToCreate := usersService.User{
		Email:    *userRequest.Email,
		Password: *userRequest.Password,
	}
	createdUser, err := h.Service.CreateUser(userToCreate)

	if err != nil {
		return nil, err
	}
	// создаем структуру респонс
	response := users.PostUsers201JSONResponse{
		Id:       &createdUser.Id,
		Email:    &createdUser.Email,
		Password: &createdUser.Password,
	}
	// Просто возвращаем респонс!
	return response, nil
}

func (h *UsersHandler) DeleteUsersId(ctx context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	// Удаляем задачу по ID
	err := h.Service.DeleteUserByID(uint(request.Id))
	if err != nil {
		// Обработайте ошибку - например, если задача не найдена
		return nil, err
	}

	// Возвращаем успешный ответ (можно вернуть пустой ответ)
	return users.DeleteUsersId204Response{}, nil
}

func (h *UsersHandler) PatchUsersId(ctx context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	// Получение данных для обновления
	updatedUser := usersService.User{}
	if request.Body.Email != nil {
		updatedUser.Email = *request.Body.Email
	}
	if request.Body.Password != nil {
		updatedUser.Password = *request.Body.Password
	}

	// Обновляем задачу
	userId := request.Id
	updatedUser, err := h.Service.UpdateUserByID(uint(userId), updatedUser)
	if err != nil {
		// Обработайте ошибку - например, если задача не найдена
		return nil, err
	}

	// Успешно обновленная задача
	response := users.PatchUsersId200JSONResponse{
		Id:       &updatedUser.Id,
		Email:    &updatedUser.Email,
		Password: &updatedUser.Password,
	}
	return response, nil
}
