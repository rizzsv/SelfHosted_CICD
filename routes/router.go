package routes

import (
	"github.com/gin-gonic/gin"
	"selfhosted-ci/handler"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.POST("/webhook", handler.WebhookHandler)
		api.GET("/build-logs", handler.GetAllLogs)
		api.GET("/build-logs/:id", handler.GetLogByID)
		api.POST("/build/manual", handler.ManualBuildHandler)
	}
}
