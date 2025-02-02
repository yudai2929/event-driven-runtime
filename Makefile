echo_build:
	go build -o functions/bin/echo functions/go/echo/main.go

logging_build:
	go build -o functions/bin/logging functions/go/logging/main.go

run_server:
	go run cmd/http/main.go

run_cron:
	go run cmd/cron/main.go

run_pubsub:
	go run cmd/pubsub/main.go

redis:
	docker run -p 6379:6379 redis