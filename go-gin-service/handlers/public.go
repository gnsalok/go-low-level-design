package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PublicEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "This is a public endpoint. Anyone can access it!",
	})
}
