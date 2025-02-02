package storage

import (
	"fmt"
	"os"
)

// FunctionStorage is an interface for storing functions.
type FunctionStorage interface {
	Exists(name string) bool
	FilePath(name string) string
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

func (s *functionStorage) Exists(name string) bool {
	filePath := s.FilePath(name)
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

func (s *functionStorage) FilePath(name string) string {
	return fmt.Sprintf("%s/%s", s.hostDir, name)
}
