package main

import (
	"go-based-splitwise/routes"
	"log"
	"net/http"
)

func main() {
	router := routes.NewRouter()
	log.Println("Server running at http://localhost:8000")
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
