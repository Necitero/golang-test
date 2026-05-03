package responses

import (
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

func SuccessfulResponse(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusAccepted, gin.H{
		"message": msg,
	})
}
