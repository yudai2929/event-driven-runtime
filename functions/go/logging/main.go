package main

import (
	"encoding/json"
	"os"

	"github.com/yudai2929/event-driven-runtime/handler"
)

const logFile = "./log.txt"

func main() {
	handler.Run(loggingHandler)
}

func loggingHandler(event map[string]any) (map[string]any, error) {
	originStdout := os.Stdout
	defer func() {
		os.Stdout = originStdout
	}()

	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	os.Stdout = file
	defer file.Close()

	_, err = file.WriteString("Hello, logging!\n")
	if err != nil {
		return nil, err
	}

	eventJSON, err := json.Marshal(event)
	if err != nil {
		return nil, err
	}

	_, err = file.Write(eventJSON)
	if err != nil {
		return nil, err
	}

	_, err = file.WriteString("\n")
	if err != nil {
		return nil, err
	}

	return event, nil
}
