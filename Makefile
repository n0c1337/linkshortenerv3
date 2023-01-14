build:
	go build cmd/linkshortener/main.go

run:
	go run cmd/linkshortener/main.go

migrate:
	go run cmd/migrator/migrator.go