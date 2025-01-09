package todo

type Service struct {
	todos []string

}

func NewService() *Service {
	return &Service{
		todos: make([]string, 0),
	}
}

func (svc *Service) AddTodo(todo string){
	svc.todos = append(svc.todos, todo)
}

func (svc *Service) GetTodos() []string{
	return svc.todos
}