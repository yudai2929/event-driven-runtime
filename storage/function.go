package storage

import (
	"fmt"
	"os"
)

// FunctionStorage is an interface for storing functions.
type FunctionStorage interface {
	Exists(name string) (bool, error)
	FilePath(name string) string
	Names() ([]string, error)
}

type functionStorage struct {
	hostDir string
}

// NewFunctionStorage creates a new function storage.
func NewFunctionStorage(hostDir string) FunctionStorage {
	return &functionStorage{
		hostDir: hostDir,
	}
}

func (s *functionStorage) Exists(name string) (bool, error) {
	filePath := s.FilePath(name)
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, fmt.Errorf("failed to stat file: %w", err)
	}

	return true, nil
}

func (s *functionStorage) FilePath(name string) string {
	return fmt.Sprintf("%s/%s", s.hostDir, name)
}

func (s *functionStorage) Names() ([]string, error) {
	dir, err := os.ReadDir(s.hostDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %w", err)
	}

	var names []string
	for _, entry := range dir {
		if entry.IsDir() {
			continue
		}
		names = append(names, entry.Name())
	}

	return names, nil
}
