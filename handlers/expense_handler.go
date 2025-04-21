package handlers

import (
	"encoding/json"
	"go-based-splitwise/models"
	"net/http"
)

var expenses = make(map[string]models.Expense) // In-memory storage for expenses

// AddExpense handles adding a new expense
func AddExpense(w http.ResponseWriter, r *http.Request) {
	var expense models.Expense
	if err := json.NewDecoder(r.Body).Decode(&expense); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	expenses[expense.ID] = expense
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Expense added successfully",
		"expense": expense,
	})
}
