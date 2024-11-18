package utils

import "github.com/gin-gonic/gin"

// SuccessResponse 成功响应
func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{"status": "success", "data": data})
}

// ErrorResponse 错误响应
func ErrorResponse(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{"status": "error", "message": message})
}
