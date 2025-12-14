package main

import (
	"github.com/gin-gonic/gin"
)

func registerRoutes(r *gin.Engine, srv *Server) {
	r.GET("/", srv.Home)
	r.GET("/auth/google", srv.GoogleLogin)
	r.GET("/auth/google/callback", srv.GoogleCallback)
	r.POST("/auth/logout", srv.Logout)
}
