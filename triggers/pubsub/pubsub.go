package pubsub

import (
	"context"
	"encoding/json"
	"log"

	"github.com/yudai2929/event-driven-runtime/pubsub"
	"github.com/yudai2929/event-driven-runtime/runtime"
	"github.com/yudai2929/event-driven-runtime/storage"
)

type subscribePayload struct {
	FunctionName string         `json:"function_name"`
	Event        map[string]any `json:"event"`
}

// StartTrigger starts the Pub/Sub trigger.
func StartTrigger(functionsDir string, addr string, chName string) {
	storage := storage.NewFunctionStorage(functionsDir)

	ctx := context.Background()
	subscriber := pubsub.NewSubscriber(addr, chName)
	subscriber.Subscribe(ctx, func(msg string) {
		var payload subscribePayload
		if err := json.Unmarshal([]byte(msg), &payload); err != nil {
			log.Fatalf("failed to unmarshal message: %v", err)
			return
		}
		filepath := storage.FilePath(payload.FunctionName)

		// Execute function
		output, err := runtime.Execute(filepath, payload.Event)
		if err != nil {
			log.Fatalf("failed to execute function: %v", err)
			return
		}

		log.Printf("Function executed with payload %s", payload)
		log.Printf("Output: %s", string(output))
	})
}
