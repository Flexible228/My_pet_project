# Таргет для создания новой миграции
migrate-new:
	migrate create -ext sql -dir ./migrations ${NAME}

# Применение миграций
migrate-up:
	migrate -path ./migrations -database "postgres://postgres:yourpassword@localhost:5432/postgres?sslmode=disable" up

# Откат миграций
migrate-down:
	migrate -path ./migrations -database "postgres://postgres:yourpassword@localhost:5432/postgres?sslmode=disable" down

genTasks:
	oapi-codegen -config openapi/.openapi.yaml -include-tags tasks -package tasks openapi/openapi.yaml > ./internal/web/tasks/api.gen.go

genUsers:
	oapi-codegen -config openapi/.openapi.yaml -include-tags users -package users openapi/openapi.yaml > ./internal/web/users/api.gen.go

lint:
	golangci-lint run --out-format=colored-line-number

 go get -u github.com/golang-migrate/migrate/v4/@v/list