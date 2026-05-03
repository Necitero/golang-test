package routes

import (
	"github.com/gin-gonic/gin"
	"necitero/golang-test/handlers"
)

func RegisterRoutes(router *gin.Engine) {
	public := router.Group("/api")
	{
		public.GET("/ping", handlers.PingHandler)
		todo := public.Group("/todo")
		{
			todo.GET("/:id", handlers.GetTodoItem)
			todo.POST("", handlers.AddTodoItem)
		}
	}
}
