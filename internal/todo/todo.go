package todo

import (
	"errors"
	"strings"
)

type Service struct {
	todos []Item

}

type Item struct {
   Task string 
   Status string
}

func NewService() *Service {
	return &Service{
		todos: make([]Item, 0),
	}
}

func (svc *Service) AddTodo(todo string) error{
	for _, t := range svc.todos{
		if t.Task == todo{
            return errors.New("todo already exists")
		}
	}
	svc.todos = append(svc.todos, Item{Task: todo, Status: "pending"})
	return nil
}

func (svc *Service) GetTodos() []Item{
	return svc.todos
}

func(svc *Service) SearchTodo(query string) []string{
	var results []string
	for _,t := range svc.todos{
		if strings.Contains(strings.ToLower(t.Task), strings.ToLower(query)){
			results = append(results, t.Task)
		}
	}
    return results
}
