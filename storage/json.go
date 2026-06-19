package storage

import (
	"encoding/json"
	"fmt"
	"go-mod-cli/todo"
	"os"
)

type JsonStorage struct {
	Filepath string
}

func (j *JsonStorage) Load(todos *[]todo.Todo) ( error) {

	data, err := os.ReadFile(j.Filepath)
	
	if err != nil {
		if os.IsNotExist(err) {
			*todos = []todo.Todo{}
			return nil
		}
		return fmt.Errorf("reading file %s : %w", j.Filepath, err)
	}

	err = json.Unmarshal(data, todos)

	if err != nil {
		return fmt.Errorf("failed conversion : %w",err)
	}

	return nil
} 

func (j *JsonStorage) Save(todos []todo.Todo) error {
	
	data, err := json.Marshal(todos)

	if err != nil {
		return fmt.Errorf("failed conversion : %w",err)
	}

	err = os.WriteFile(j.Filepath, data, 0644)

	if err != nil {
		return fmt.Errorf("Failed to write : %w", err)
	}

	return nil
}
