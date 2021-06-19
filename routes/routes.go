package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/insomniacoder/go-kafka/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.POST("comments", controllers.PublishComment)
	}
	return r
}
