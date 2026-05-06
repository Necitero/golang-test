package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"necitero/golang-test/database"
	"necitero/golang-test/models"
	"necitero/golang-test/responses"
	"strconv"
)

func GetTodoItem(ctx *gin.Context) {
	db, err := database.OpenDatabase()
	if err != nil {
		responses.CouldNotRetrieveDB(ctx)
	}
	id := ctx.Param("id")
	var entry *models.DBEntry
	for i := 0; i < len(db.Todos); i++ {
		index := db.Todos[i].Index
		if strconv.Itoa(index) == id {
			entry = &db.Todos[i]
		}
	}
	if entry == nil {
		responses.CouldNotFindEntry(ctx)
	} else {
		responses.SuccessfulToDo(ctx, entry)
	}
}

func AddTodoItem(ctx *gin.Context) {
	db, err := database.OpenDatabase()
	if err != nil {
		responses.CouldNotRetrieveDB(ctx)
	}
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
	fmt.Println("DATA", todo.Headline, todo.Note, todo.Status, todo.DueDate)
	database.UpdateDatabase(db, &todo)
	responses.SuccessfulResponse(ctx, "ToDo item added")
}
