package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RequestLogger(logger *zap.SugaredLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()

		logger.Infow("incoming request",
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"status", status,
			"duration", latency,
		)
	}
}
