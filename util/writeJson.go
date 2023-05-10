package util

import "github.com/gin-gonic/gin"

func WriteJSON(c *gin.Context, status int, v any) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(status)

	c.JSONP(status, gin.H{
		"status": status,
		"data":   v,
	})
}
