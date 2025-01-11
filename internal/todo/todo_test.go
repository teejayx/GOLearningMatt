package todo_test

import (
	"my-first-api/internal/todo"
	"reflect"
	"testing"
)

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
		  got := svc.SearchTodo(tt.query)
		  if !reflect.DeepEqual(got, tt.expectedResult){
			  t.Errorf("SearchTodo() = %v, want %v", got, tt.expectedResult)
		  }
	  })
  }
    

}