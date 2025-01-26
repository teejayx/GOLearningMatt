package main

import (
	"log"
	db "my-first-api/internal/Db"
	"my-first-api/internal/todo"
	"my-first-api/internal/transport"
)



func main() {
     
	d, err :=  db.New("root","root", "ny_taxi","localhost", 5432)	
	if err != nil{
		log.Fatal(err)
	}
	svc := todo.NewService(d)
	server := transport.NewServer(svc)
 
	if err := server.Serve(); err != nil {
		log.Fatal(err)
	}



	log.Println("Server starting on :8080")
	
}