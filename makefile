# Таргет для создания новой миграции
migrate-new:
	migrate create -ext sql -dir ./migrations add_user_id
# Применение миграций
migrate-up:
	migrate -path ./migrations -database "postgres://postgres:yourpassword@localhost:5432/postgres?sslmode=disable" up

# Откат миграций
migrate-down:
	migrate -path ./migrations -database "postgres://postgres:yourpassword@localhost:5432/postgres?sslmode=disable" down

# Генерация тега tasks
genTasks:
	oapi-codegen -config openapi/.openapi.yaml -include-tags tasks -package tasks openapi/openapi.yaml > ./internal/web/tasks/api.gen.go

# Генерация тега users
genUsers:
	oapi-codegen -config openapi/.openapi.yaml -include-tags users -package users openapi/openapi.yaml > ./internal/web/users/api.gen.go

# Линтер для аналаиза кодовой базы на наличие ошибок
lint:
	golangci-lint run --out-format=colored-line-number
