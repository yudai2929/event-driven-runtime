package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func executeBinary(binaryPath string) (string, error) {
	cmd := exec.Command(binaryPath)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("execution failed: %v, stderr: %s", err, stderr.String())
	}

	return stdout.String(), nil
}

func main() {
	binaryPath := "./bin/hello" // 必要に応じてパスを設定

	output, err := executeBinary(binaryPath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Output: %s\n", output)
}
