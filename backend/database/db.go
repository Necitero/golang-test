package database

import (
	"encoding/json"
	"necitero/golang-test/models"
	"os"
)

var DATABASE_PATH string = "./database/database.json"

func OpenDatabase() (*models.DBData, error) {
	var data models.DBData
	content, err := os.ReadFile(DATABASE_PATH)
	if err != nil {
		return &data, err
	}

	if err := json.Unmarshal(content, &data); err != nil {
		return &data, err
	}
	return &data, nil
}

func UpdateDatabase(db *models.DBData, todo *models.ToDo) {
	var entry models.DBEntry
	entry.Todo = *todo
	entry.Index = len(db.Todos)
	db.Todos = append(db.Todos, entry)
	bytes, err := json.MarshalIndent(db, "", "  ")
	if err != nil {
		return
	}
	os.WriteFile(DATABASE_PATH, bytes, 0644)
}
