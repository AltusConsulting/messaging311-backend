package models

type Message struct {
	To             string       `json:"to" binding:"required"`
	DelayWhileIdle bool         `json:"delay_while_idle" binding:"required"`
	Priority       string       `json:"priority" binding:"required"`
	Data           Data         `json:"data" binding:"required"`
	Notification   Notification `json:"notification" binding:"required"`
}
