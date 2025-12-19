package main

import "github.com/gin-gonic/gin"

func secureHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set(
			"Content-Security-Policy",
			"default-src 'self'; "+
				"img-src 'self' https://lh3.googleusercontent.com;"+
				"style-src 'self' 'unsafe-inline' https://fonts.googleapis.com; "+
				"font-src 'self' https://fonts.gstatic.com",
		)
		c.Next()
	}
}
