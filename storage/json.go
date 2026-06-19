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
		return fmt.Errorf("reading file %s : %w", j.Filepath, err)
	}

	data_err := json.Unmarshal(data, todos)

	if data_err != nil {
		return fmt.Errorf("failed conversion : %w",data_err)
	}

	return nil
} 

func (j *JsonStorage) Save(todos []todo.Todo) error {
	
	byte_data, err := json.Marshal(todos)

	if err != nil {
		return fmt.Errorf("failed conversion : %w",err)
	}

	err = os.WriteFile(j.Filepath, byte_data, 0777)

	if err != nil {
		return fmt.Errorf("Failed to write : %w", err)
	}

	return nil
}
