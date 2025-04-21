package storage

import "go-based-splitwise/models"

var Users = make(map[string]models.User)
var Expenses = make(map[string]models.Expense)
var Groups = make(map[string]models.Group)
