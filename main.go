package main

import (
	"encoding/json"
	"log"
	"my-first-api/internal/todo"
	"net/http"
)

type TodoItem struct{
	Item string `json:"item"`
} 


func main() {
	svc := todo.NewService()
	mux := http.NewServeMux()
 
	mux.HandleFunc("GET /todo", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		response, err := json.Marshal(svc.GetTodos())
		if err != nil {
			http.Error(w, "Error converting todos to JSON", http.StatusInternalServerError)
			return
		}
		_, err = w.Write(response)
		if err != nil {
			log.Printf("Error writing response: %v", err)
		}
	})

	mux.HandleFunc("POST /todo", func(w http.ResponseWriter, r *http.Request) {	
		var t TodoItem
		err := json.NewDecoder(r.Body).Decode(&t)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		svc.AddTodo(t.Item)
		w.WriteHeader(http.StatusCreated)
		return
        
	})





	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}