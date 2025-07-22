package config

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupCORS(r *gin.Engine) {
	r.Use(cors.Default())
}
