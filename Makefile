echo_build:
	go build -o functions/bin/echo functions/go/echo/main.go

logging_build:
	go build -o functions/bin/logging functions/go/logging/main.go

run_server:
	go run cmd/http/main.go

