package runtime

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
)

// Execute is a function that executes a function with the given name and event.
func Execute(functionPath string, event map[string]any) ([]byte, error) {

	eventJSON, err := json.Marshal(event)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal event: %w", err)
	}

	cmd := exec.Command(functionPath)
	cmd.Stdin = bytes.NewReader(eventJSON)

	var output bytes.Buffer
	cmd.Stdout = &output
	cmd.Stderr = &output

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("failed to execute function: %w", err)
	}

	if cmd.ProcessState.ExitCode() != 0 {
		return nil, fmt.Errorf("function failed with exit code %d: %s", cmd.ProcessState.ExitCode(), output.String())
	}

	return output.Bytes(), nil
}
