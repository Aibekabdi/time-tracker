package utils

import "github.com/gin-gonic/gin"

func ErrorHandler(c *gin.Context, status int, msg string) {
	c.JSON(status, map[string]string{
		"error": msg,
	})
}
