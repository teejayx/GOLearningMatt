package todo_test

import (
	"context"
	db "my-first-api/internal/Db"
	"my-first-api/internal/todo"
	"reflect"
	"testing"
)

type MockDB struct{
items []db.Item
}

func (m *MockDB) InsertItems (ctx context.Context, item db.Item) error {
	m.items = append(m.items, item)
	return nil
}


func (m *MockDB) GetAllItems(ctx context.Context, item db.Item) ([]db.Item, error) {
	return m.items, nil
}



func TestService_Service(t *testing.T){


	tests := []struct{
		name string
		query string
		toDosToAdd []string
		expectedResult []string
	}{
       { name: "Search for 'task'", 
	    query: "sh", 
		toDosToAdd: []string{"shop"}, 
		expectedResult: []string{"shop"}},
	}
  
  for _, tt := range tests{
	  t.Run(tt.query, func(t *testing.T){
		  svc := todo.NewService()
		  for _, toAdd := range tt.toDosToAdd{
			  err := svc.AddTodo(toAdd)
			  if err != nil{
				  t.Error(err)
			  }
		  }
		  got,err := svc.SearchTodo(tt.query)
		  if err != nil{
			  t.Error(err)
		  }
		  if !reflect.DeepEqual(got, tt.expectedResult){
			  t.Errorf("SearchTodo() = %v, want %v", got, tt.expectedResult)
		  }
	  })
  }
    

}