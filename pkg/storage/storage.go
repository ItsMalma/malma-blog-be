package storage

import (
	"os"
	"path/filepath"
)

type StorageManager struct {
	baseDirectory string
}

func NewStorageManager(baseDirectory string) StorageManager {
	return StorageManager{baseDirectory: baseDirectory}
}

func (storageManager StorageManager) Load(fileName string) ([]byte, error) {
	return os.ReadFile(filepath.Join(storageManager.baseDirectory, fileName))
}

// Put the fileName with extension!
func (storageManager StorageManager) Save(fileName string, content []byte) error {
	return os.WriteFile(filepath.Join(storageManager.baseDirectory, fileName), content, 0777)
}

func (storageManager StorageManager) Delete(fileName string) error {
	return os.Remove(filepath.Join(storageManager.baseDirectory, fileName))
}
