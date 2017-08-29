package models

type Notification struct {
	Body  string `json:"body" binding:"required"`
	Title string `json:"title" binding:"required"`
	Sound string `json:"sound" binding:"required"`
}
