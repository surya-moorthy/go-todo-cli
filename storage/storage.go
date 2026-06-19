package storage

import "go-mod-cli/todo"

type Storage interface {
	Load() ([]todo.Todo, error)
	Save([]todo.Todo) error
}