package storage

import "go-mod-cli/model"

type Storage interface {
	Load() ([]model.Todo, error)
	Save([]model.Todo) error
}