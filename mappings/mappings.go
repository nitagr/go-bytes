package mappings

import (
	"geektrust/handlers"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func CreateUrlMappings() {
	Router = gin.New()
	api := Router.Group("/api")
	{
		api.GET("/test", handlers.SendSuccessmessage)
		api.GET("/sum", handlers.CalculateSum)
	}
}
