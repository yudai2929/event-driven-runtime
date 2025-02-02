package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yudai2929/event-driven-runtime/storage"
)

// StartTrigger is a function that starts the trigger
func StartTrigger(functionsDir string, port int) error {
	mux := http.NewServeMux()

	// di
	storage := storage.NewFunctionStorage(functionsDir)
	h := &handler{
		storage: storage,
	}

	// routes
	mux.HandleFunc("POST /invoke", loggingMiddleware(h.invokeHandler))
	mux.HandleFunc("GET /functions/{function_name}", loggingMiddleware(h.getFunctionHandler))
	mux.HandleFunc("GET /functions", loggingMiddleware(h.listFunctionHandler))

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	log.Printf("Server listening on %s", server.Addr)
	log.Printf("http://localhost%s/invoke", server.Addr)

	return server.ListenAndServe()
}
