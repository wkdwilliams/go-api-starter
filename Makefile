run:
	@nodemon --watch './**/*.go' --signal SIGTERM --exec 'clear && go' run cmd/server/main.go
build:
	@go build -o bin/out cmd/server/main.go
migrate-sql:
	@go run cmd/migrate/main.go
swagger:
	@swag init --parseDependency --parseInternal -d cmd/server/,internal/api --output docs/
test:
	@go test internal/service/* 