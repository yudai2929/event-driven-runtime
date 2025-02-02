package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/yudai2929/event-driven-runtime/runtime"
	"github.com/yudai2929/event-driven-runtime/storage"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /invoke", loggingMiddleware(invokeHandler))
	mux.HandleFunc("GET /functions/{function_name}", loggingMiddleware(getFunctionHandler))
	mux.HandleFunc("GET /functions", loggingMiddleware(listFunctionHandler))

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Printf("Server listening on %s", server.Addr)
	log.Printf("http://localhost%s/invoke", server.Addr)
	log.Fatal(server.ListenAndServe())
}

const functionsDir = "functions/bin"

// InvokeRequest is the payload of the request.
type InvokeRequest struct {
	FunctionName string         `json:"function_name"`
	Event        map[string]any `json:"event"`
}

func invokeHandler(w http.ResponseWriter, r *http.Request) {

	var payload InvokeRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "failed to decode request payload", http.StatusBadRequest)
		return
	}

	storage := storage.NewFunctionStorage(functionsDir)
	ok, err := storage.Exists(payload.FunctionName)
	if err != nil {
		http.Error(w, "failed to check function existence", http.StatusInternalServerError)
		return
	}
	if !ok {
		http.Error(w, "function not found", http.StatusNotFound)
		return
	}

	functionsPath := storage.FilePath(payload.FunctionName)
	output, err := runtime.Execute(functionsPath, payload.Event)
	if err != nil {
		http.Error(w, "failed to execute function", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

// GetFunctionResponse is the payload of the response.
type GetFunctionResponse struct {
	FunctionName string `json:"function_name"`
}

func getFunctionHandler(w http.ResponseWriter, r *http.Request) {
	functionName := r.PathValue("function_name")
	if functionName == "" {
		http.Error(w, "function_name is required", http.StatusBadRequest)
		return
	}

	storage := storage.NewFunctionStorage(functionsDir)
	ok, err := storage.Exists(functionName)
	if err != nil {
		http.Error(w, "failed to check function existence", http.StatusInternalServerError)
		return
	}
	if !ok {
		http.Error(w, "function not found", http.StatusNotFound)
		return
	}

	response := &GetFunctionResponse{
		FunctionName: functionName,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "failed to encode response payload", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}

func listFunctionHandler(w http.ResponseWriter, _ *http.Request) {
	storage := storage.NewFunctionStorage(functionsDir)
	names, err := storage.Names()
	if err != nil {
		http.Error(w, "failed to list functions", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(names); err != nil {
		http.Error(w, "failed to encode response payload", http.StatusInternalServerError)
		return
	}
}

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL)
		next(w, r)
	}
}
