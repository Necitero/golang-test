package models

import "time"

type ToDo struct {
	Headline     string     `json:"headline" binding:"required"`
	Note         string     `json:"note,omitempty"`
	Status       string     `json:"status,omitempty"`
	UrgencyLevel int        `json:"urgency_level,omitempty"`
	DueDate      *time.Time `json:"due_date,omitempty"`
}

type DBEntry struct {
	Todo  ToDo `json:"todo" binding:"required"`
	Index int  `json:"id" binding:"required"`
}

type DBData struct {
	Todos []DBEntry `json:"todos" binding:"required"`
}
