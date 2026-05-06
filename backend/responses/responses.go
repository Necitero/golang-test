package responses

import (
	"necitero/golang-test/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InvalidContentType(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"message": "Invalid content type",
	})
}

func InvalidBodyContent(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"message": "Invalid body content",
	})
}

func SuccessfulToDo(ctx *gin.Context, todo *models.DBEntry) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": todo,
	})
}

func SuccessfulResponse(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusAccepted, gin.H{
		"message": msg,
	})
}

func CouldNotRetrieveDB(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"message": "Could not retrieve database",
	})
}

func CouldNotFindEntry(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, gin.H{
		"message": "Could not find entry in database",
	})
}
