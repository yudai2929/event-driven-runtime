package handler

import (
	"encoding/json"
	"fmt"
	"os"
)

// HandleFunc is a function that takes an event and returns a response.
type HandleFunc func(event map[string]any) (map[string]any, error)

// Run is a function that runs the handler.
func Run(handler HandleFunc) {
	var event map[string]any
	if err := json.NewDecoder(os.Stdin).Decode(&event); err != nil {
		fmt.Fprintf(os.Stderr, "failed to decode event: %v", err)
		os.Exit(1)
	}

	response, err := handler(event)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to handle event: %v", err)
		os.Exit(1)
	}

	if err := json.NewEncoder(os.Stdout).Encode(response); err != nil {
		fmt.Fprintf(os.Stderr, "failed to encode response: %v", err)
		os.Exit(1)
	}

	os.Exit(0)
}
