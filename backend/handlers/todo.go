package handlers

import (
	"necitero/golang-test/database"
	"necitero/golang-test/models"
	"necitero/golang-test/responses"
	"necitero/golang-test/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetTodoItem(ctx *gin.Context) {
	id := ctx.Param("id")
	query := "SELECT * FROM " + database.TableName + " WHERE id = " + id + " LIMIT 1"
	utils.Logger("TODO", query)
	data := database.GetRow(query)
	var todo models.DatabaseEntry
	err := data.Scan(&todo.Id, &todo.Headline, &todo.Note, &todo.Status, &todo.UrgencyLevel, &todo.DueDate)
	if err != nil {
		responses.CouldNotFindEntry(ctx)
		panic(err)
	}
	responses.SuccessfulToDo(ctx, &todo)
}

func AddTodoItem(ctx *gin.Context) {
	header := ctx.Request.Header
	requestType := header.Get("Content-Type")
	if requestType != "application/json" {
		responses.InvalidContentType(ctx)
		return
	}

	var todo models.ToDo
	if err := ctx.ShouldBindJSON(&todo); err != nil {
		responses.InvalidBodyContent(ctx)
		return
	}
	if todo.Status == "" {
		todo.Status = "not_started"
	}

	// Set insertion query
	query := "INSERT INTO " + database.TableName + " (headline, note, status, urgency, due_date)"
	query += " VALUES( "
	query += utils.DBStringInsert(todo.Headline, false)
	query += utils.DBStringInsert(todo.Note, false)
	query += utils.DBStringInsert(todo.Status, false)
	query += strconv.Itoa(todo.UrgencyLevel) + ","
	if todo.DueDate != nil {
		dateStr := todo.DueDate.Format("2006-01-02 15:04:05")
		query += utils.DBStringInsert(dateStr, true)
	} else {
		query += "NULL"
	}
	query += " )"
	utils.Logger("DB", "Inserting with query: "+query)

	_, success := database.ExecQuery(query)
	if !success {
		responses.QueryFailed(ctx, query)
		return
	}
	responses.SuccessfulResponse(ctx, "Created new ToDo: "+todo.Headline)
}
