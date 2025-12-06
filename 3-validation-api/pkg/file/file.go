package file

import (
	json2 "encoding/json"
	"fmt"
	"github.com/ANB98prog/purple-school-homeworks/3-validation-api/pkg/json"
	"os"
)

func ReadFile[T any](filePath string) (*T, error) {
	file, err := os.OpenFile(filePath, os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileData, err := json.Decode[T](file)
	if err != nil {
		return nil, fmt.Errorf("file %s is not a valid JSON file. Error: %e", filePath, err)
	}

	err = json.IsValid(fileData)
	if err != nil {
		return nil, fmt.Errorf("file %s is not a valid JSON file. Error: %e", filePath, err)
	}

	return &fileData, nil
}

func WriteFile[T any](filePath string, data *T) error {
	file, err := os.OpenFile(filePath, os.O_CREATE, os.ModePerm)
	if err != nil {
		return fmt.Errorf("could not write to file %s. Error: %e", filePath, err)
	}
	json2.NewEncoder()
	os.WriteFile(filePath)
}
