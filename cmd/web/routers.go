package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func registerRoutes(r *gin.Engine, srv *Server) {
	r.GET("/", srv.Home)
	r.GET("/auth/google", srv.HandleGoogleLogin)
	r.GET("/auth/google/callback", srv.HandleGoogleCallback)
	r.POST("/auth/logout", srv.AuthRequired(), srv.HandleLogout)
	r.GET("/auth/logout", srv.AuthRequired(), srv.HandleLogout)
	r.GET("/profile", srv.AuthRequired(), srv.HandleProfile)
	r.GET("/test-flash", func(c *gin.Context) {
		srv.ErrorMessage(c, "This is an error message")
		srv.InfoMessage(c, "This is an info message")
		srv.SuccessMessage(c, "This is a success message")

		c.Redirect(http.StatusSeeOther, "/")
	})
}
