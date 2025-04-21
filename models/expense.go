package models

type Expense struct {
	ID           string   `json:"id"`
	Amount       float64  `json:"amount"`
	PayerID      string   `json:"payer_id"`
	SplitBetween []string `json:"split_between"` // user IDs of people in the group
	Description  string   `json:"description"`
}
