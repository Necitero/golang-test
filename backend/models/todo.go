package models

import "time"

type ToDo struct {
	Headline     string     `json:"headline" binding:"required"`
	Note         string     `json:"note,omitempty"`
	Status       string     `json:"status,omitempty"`
	UrgencyLevel int        `json:"urgency_level,omitempty"`
	DueDate      *time.Time `json:"due_date,omitempty"`
}
