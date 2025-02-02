package main

import (
	"fmt"

	"github.com/yudai2929/event-driven-runtime/handler"
)

func main() {
	handler.Run(echoHandler)
}

func echoHandler(event map[string]any) (map[string]any, error) {
	name, ok := event["name"].(string)
	if !ok {
		name = "World"
	}

	return map[string]any{
		"message": fmt.Sprintf("Hello %s!", name),
	}, nil
}
