package handler

import "github.com/gin-gonic/gin"

// GET /ping
func PingGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	}
}
