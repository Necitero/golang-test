package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"necitero/golang-test/utils"
)

var TableName string = "todos"

func getDatabase() *sql.DB {
	var connection string = "app_user:app_password@tcp(db:3306)/app_db?parseTime=true"
	db, err := sql.Open("mysql", connection)
	if err != nil {
		utils.Logger("DB", "Error opening database. Panicking.")
		panic(err)
	} else if err = db.Ping(); err != nil {
		utils.Logger("DB", "Error pinging database. Panicking.")
		panic(err)
	}
	return db
}

func GetRow(query string) *sql.Row {
	db := getDatabase()
	data := db.QueryRow(query)
	return data
}

func ExecQuery(query string) (sql.Result, bool) {
	db := getDatabase()
	data, err := db.Exec(query)
	defer db.Close()
	if err != nil {
		msg := "Error executing query: " + query
		utils.Logger("DB", msg)
		return data, false
	}
	return data, true
}

func SetupDatabase() {
	// Table create string
	var tableCreate string = "id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY"
	tableCreate += ", headline TEXT NOT NULL"
	tableCreate += ", note TEXT"
	tableCreate += ", status TEXT NOT NULL"
	tableCreate += ", urgency INT NOT NULL"
	tableCreate += ", due_date DATETIME"

	var instruction string = "CREATE TABLE IF NOT EXISTS " + TableName + " (" + tableCreate + ")"
	_, success := ExecQuery(instruction)

	if success {
		utils.Logger("DB", "Successfully created database \"todos\"")
	}
}
