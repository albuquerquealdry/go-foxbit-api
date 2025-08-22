package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Liveness(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "alive"})
}

func Readiness(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ready"})
}
