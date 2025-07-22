package handler

import (
	"net/http"
	"selfhosted-ci/storage"
	"github.com/gin-gonic/gin"
)

func GetAllLogs(c *gin.Context) {
	logs := storage.GetAllLogs()
	c.JSON(http.StatusOK, logs)
}

func GetLogByID(c *gin.Context) {
	id := c.Param("id")
	log := storage.GetLogByID(id)

	if log == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "log not found"})
		return
	}

	c.JSON(http.StatusOK, log)
}

func ManualBuildHandler(c *gin.Context) {
	var payload struct {
		Repo string `json:"repo"`
	}
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	log := storage.AddLog(payload.Repo, "manual_triggered", "Manual build triggered successfully")
	c.JSON(http.StatusOK, log)
}
