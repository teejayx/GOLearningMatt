package todo

import (
	"context"
	"errors"
	"fmt"
	db "my-first-api/internal/Db"
	"strings"
)

type Service struct {
	db Manager

}

type Item struct {
   Task string 
   Status string
}

type Manager interface {
	InsertItems(ctx context.Context, item db.Item) error
	GetItems(ctx context.Context) ([]db.Item, error)
}



func NewService(db Manager) *Service {
	return &Service{
		db: db,
	}
}

func (svc *Service) AddTodo(todo string) error{
    items, err := svc.GetAll()
	if err != nil{
		return fmt.Errorf("error getting items from db: %v", err)
	}


	for _, t := range items{
		if t.Task == todo{
            return errors.New("todo already exists")
		}
	}
	if err := svc.db.InsertItems(context.Background(), db.Item{Task: todo, Status: "pending"}); err != nil{
		return fmt.Errorf("error inserting item into db: %v", err)
	}
	return nil
}

func(svc *Service) SearchTodo(query string) ([]string, error){
	items, err := svc.GetAll()
	if err != nil{
		return nil, fmt.Errorf("error getting items from db: %v", err)
	}
	var results []string
	for _,t := range items{
		if strings.Contains(strings.ToLower(t.Task), strings.ToLower(query)){
			results = append(results, t.Task)
		}
	}
    return results, nil
}

func (svc *Service) GetAll() ([]Item, error){
	var results []Item
	items, err := svc.db.GetItems(context.Background())
	if err != nil{
		return nil, fmt.Errorf("error getting items from db: %v", err)	
	}

	for _, item := range items{
		results = append(results, Item{Task: item.Task, Status: item.Status})
	}
    return results, nil
}
