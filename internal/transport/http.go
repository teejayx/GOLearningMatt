package transport

import (
	"encoding/json"
	"log"
	"my-first-api/internal/todo"
	"net/http"
)

type TodoItem struct{
	Item string `json:"item"`
} 

type Server struct {
	mux *http.ServeMux
}

func NewServer(todoSvc *todo.Service) *Server {

     mux := http.NewServeMux()

	mux.HandleFunc("GET /todo", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		response, err := json.Marshal(todoSvc.GetTodos())
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

		err = todoSvc.AddTodo(t.Item)

		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
		return
        
	})
	return &Server{mux: mux}


}

func (s *Server) Serve() error {
	return http.ListenAndServe(":8080", s.mux)
}