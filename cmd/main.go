package main

import (
	"fmt"

	pkgruntime "github.com/yudai2929/event-driven-runtime/runtime"
)

func main() {
	runtime := pkgruntime.NewRuntime(10)

	runtime.RegisterHandler("greet", func(event pkgruntime.Event) string {
		return fmt.Sprintf("Hello, Event ID: %d!", event.ID)
	})
	runtime.RegisterHandler("goodbye", func(event pkgruntime.Event) string {
		return fmt.Sprintf("Goodbye, Event ID: %d!", event.ID)
	})

	runtime.Start()

	go func() {
		runtime.Emit(pkgruntime.Event{ID: 1, Payload: "greet"})
		runtime.Emit(pkgruntime.Event{ID: 2, Payload: "goodbye"})
		runtime.Emit(pkgruntime.Event{ID: 3, Payload: "unknown"})

		runtime.Stop()
	}()

	for response := range runtime.ListenResponses() {
		fmt.Println(response)
	}
}
