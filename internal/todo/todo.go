package todo

import "errors"

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