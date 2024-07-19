run:
	@nodemon --watch './**/*.go' --signal SIGTERM --exec 'clear && go' run cmd/server/server.go
build:
	@go build -o bin/out cmd/server/server.go
migrate-sql:
	@go run cmd/migrate/migrate_sql.go
test:
	@go test ./tests