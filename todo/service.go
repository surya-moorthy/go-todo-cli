package todo

import (
	"fmt"
	"go-mod-cli/model"
	"go-mod-cli/storage"
)

type Service struct {
	Todos []model.Todo
	Storage storage.Storage
}

func (s *Service) Add(t model.Todo) (string, error) {
	s.Todos = append(s.Todos, t)
	err := s.Storage.Save(s.Todos)
	
	if err != nil {
		return "", fmt.Errorf("%s", err.Error())
	}

	return "Added successfully", nil
}

func (s *Service) Delete(title string) (string, error) {
	found := false

	for i,todo := range s.Todos {
		if todo.Title == title {
			s.Todos = append(s.Todos[:i] , s.Todos[i+1:]...)
			found = true
		}
	}

	if !found {
		return "", fmt.Errorf("todo not found.")
	}

	err := s.Storage.Save(s.Todos)

	if err != nil {
		return "", fmt.Errorf("%s", err.Error())
	}

	return "successfully deleted",nil
}

func (s *Service) Update(title string, t model.Todo) (string, error) {
	found := false

	for i,todo := range s.Todos {
		if todo.Title == title {
			s.Todos[i].Title = t.Title
			s.Todos[i].Description = t.Description
			found = true
		}
	}

	if !found {
		return "", fmt.Errorf("todo not found.")
	}

	err := s.Storage.Save(s.Todos)

	if err != nil {
		return "", fmt.Errorf("%s", err.Error())
	}

	return "updated successfully", nil
}

func (s *Service) List() ([]model.Todo, error) {
	return s.Todos, nil
}

func (s *Service) Search(title string) (model.Todo, error) {
	var todo model.Todo
	found := false

	for _,t := range s.Todos {
		if t.Title == title {
			todo = t
			found = true
			break
		}
	}

	if !found {
		return todo, fmt.Errorf("todo not found.")
	}

	return todo, nil
}