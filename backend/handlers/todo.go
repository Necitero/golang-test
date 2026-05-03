package handlers

import (
	"fmt"
	"necitero/golang-test/models"
	"necitero/golang-test/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTodoItem(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, gin.H{
		"message": "Aaaah nooo so sad, no info :<",
	})
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
		fmt.Println(err)
		responses.InvalidBodyContent(ctx)
		return
	}
	fmt.Println("DATA", todo.Headline, todo.Note, todo.Status, todo.DueDate)
	responses.SuccessfulResponse(ctx, "ToDo item added")
}
