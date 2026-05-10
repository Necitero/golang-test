package models

import "time"

type DatabaseEntry struct {
	Id           int
	Headline     string
	Note         string
	Status       string
	UrgencyLevel int
	DueDate      *time.Time
}
