package main

import "github.com/yudai2929/event-driven-runtime/triggers/http"

const (
	functionsDir = "functions/bin"
	port         = 8080
)

func main() {
	http.StartTrigger(functionsDir, port)
}
