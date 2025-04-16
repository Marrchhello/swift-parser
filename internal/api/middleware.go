package api

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger middleware logs request details
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Process request
		c.Next()

		// Log details after request is processed
		log.Printf(
			"[%s] %s %s %v",
			c.Request.Method,
			c.Request.URL.Path,
			c.ClientIP(),
			time.Since(start),
		)
	}
}

// ErrorHandler middleware catches panics
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic recovered: %v", err)
				c.JSON(500, gin.H{
					"error": "Internal server error",
				})
			}
		}()
		c.Next()
	}
}

// ValidateSwiftCode middleware validates SWIFT code format
func ValidateSwiftCode() gin.HandlerFunc {
	return func(c *gin.Context) {
		if swiftCode := c.Param("swiftCode"); len(swiftCode) != 11 {
			c.JSON(400, gin.H{"error": "Invalid SWIFT code format"})
			c.Abort()
			return
		}
		c.Next()
	}
}
