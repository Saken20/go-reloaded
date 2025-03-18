package fileio

import (
	"io/ioutil"
)

// Чтение файла
func ReadFile(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// Запись файла
func WriteFile(filename string, content string) error {
	return ioutil.WriteFile(filename, []byte(content), 0644)
}
