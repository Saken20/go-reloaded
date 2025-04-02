package fileio

import (
	"os"
)

// Чтение файла
func ReadFile(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// Запись файла
func WriteFile(filename string, content string) error {
	return os.WriteFile(filename, []byte(content), 0644)
}
