package handlers

import (
	"encoding/json"
	"go-based-splitwise/models"
	storage "go-based-splitwise/storages"
	"net/http"
)

// CreateGroup creates a new group
func CreateGroup(w http.ResponseWriter, r *http.Request) {
	var group models.Group
	if err := json.NewDecoder(r.Body).Decode(&group); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	storage.Groups[group.ID] = group

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Group created successfully",
		"groupId": group.ID,
	})
}
