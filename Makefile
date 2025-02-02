echo_build:
	go build -o bin/echo functions/echo/main.go

logging_build:
	go build -o bin/logging functions/logging/main.go

run_server:
	go run cmd/http/main.go

