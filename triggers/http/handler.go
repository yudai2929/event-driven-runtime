package http

import (
	"encoding/json"
	"net/http"

	"github.com/yudai2929/event-driven-runtime/runtime"
	"github.com/yudai2929/event-driven-runtime/storage"
)

type handler struct {
	storage storage.FunctionStorageClient
}

// InvokeRequest is the payload of the request.
type InvokeRequest struct {
	FunctionName string         `json:"function_name"`
	Event        map[string]any `json:"event"`
}

func (h *handler) invokeHandler(w http.ResponseWriter, r *http.Request) {
	var payload InvokeRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "failed to decode request payload", http.StatusBadRequest)
		return
	}

	ok, err := h.storage.Exists(payload.FunctionName)
	if err != nil {
		http.Error(w, "failed to check function existence", http.StatusInternalServerError)
		return
	}
	if !ok {
		http.Error(w, "function not found", http.StatusNotFound)
		return
	}

	functionsPath := h.storage.FilePath(payload.FunctionName)
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

func (h *handler) getFunctionHandler(w http.ResponseWriter, r *http.Request) {
	functionName := r.PathValue("function_name")
	if functionName == "" {
		http.Error(w, "function_name is required", http.StatusBadRequest)
		return
	}

	ok, err := h.storage.Exists(functionName)
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

func (h *handler) listFunctionHandler(w http.ResponseWriter, _ *http.Request) {
	names, err := h.storage.Names()
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
