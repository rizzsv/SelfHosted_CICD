package handler

import (
	"net/http"
	"selfhosted-ci/storage"
	"github.com/gin-gonic/gin"
)

type WebhookPayload struct {
	Repository struct {
		Name string `json:"name"`
	} `json:"repository"`
}

func WebhookHandler(c *gin.Context) {
	var payload WebhookPayload

	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	log := storage.AddLog(payload.Repository.Name, "success", "Webhook received and build triggered")
	c.JSON(http.StatusOK, log)
}
