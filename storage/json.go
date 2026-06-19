package storage

import (
	"encoding/json"
	"fmt"
	"go-mod-cli/model"
	"os"
)

type JsonStorage struct {
	Filepath string
}

func (j *JsonStorage) Load() ([]model.Todo, error) {

	data, err := os.ReadFile(j.Filepath)

	if err != nil {
		if os.IsNotExist(err) {
			return []model.Todo{}, nil
		}
		return nil,fmt.Errorf("reading file %s : %w", j.Filepath, err)
	}

	if len(data) == 0 {
		return []model.Todo{}, nil
	}
	var todos []model.Todo

	err = json.Unmarshal(data, &todos)

	if err != nil {
		return nil,fmt.Errorf("failed conversion : %w", err)
	}

	return todos,nil
}

func (j *JsonStorage) Save(todos []model.Todo) error {

	data, err := json.Marshal(todos)

	if err != nil {
		return fmt.Errorf("failed conversion : %w", err)
	}

	err = os.WriteFile(j.Filepath, data, 0644)

	if err != nil {
		return fmt.Errorf("Failed to write : %w", err)
	}

	return nil
}
