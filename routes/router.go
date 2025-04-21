package routes

import (
	"go-based-splitwise/handlers"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	// User routes
	router.HandleFunc("/api/users/register", handlers.RegisterUser).Methods("POST")

	// Expense routes
	router.HandleFunc("/api/expenses", handlers.AddExpense).Methods("POST")

	// Group routes
	router.HandleFunc("/api/groups", handlers.CreateGroup).Methods("POST")

	return router
}
