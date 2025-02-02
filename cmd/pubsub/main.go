package main

import "github.com/yudai2929/event-driven-runtime/triggers/pubsub"

const (
	functionsDir = "functions/bin"
	addr         = "localhost:6379"
	chName       = "pubsub"
)

func main() {
	// Start the Pub/Sub trigger
	pubsub.StartTrigger(functionsDir, addr, chName)
}
