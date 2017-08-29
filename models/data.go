package models

type Data struct {
	Alert bool   `json:"alert" binding:"required"`
	ID    string `json:"id" binding:"required"`
	Page  string `json:"page" binding:"required"`
}
