package main

import (
	"context"
	"fmt"
)

type MyEvent struct {
	Name string `json:"name"`
}

type MyResponse struct {
	Message string `json:"message"`
}

func HandleRequest(ctx context.Context, name MyEvent) (MyResponse, error) {
	return MyResponse{Message: fmt.Sprintf("Hello %s!", name.Name)}, nil
}

func main() {
	// lambda.Start(HandleRequest)

}
