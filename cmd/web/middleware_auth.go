package main

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/osamah22/open-mart/internal/services"
)

const ContextUserKey = "user"

func AuthContext(authSvc services.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		sess := sessions.Default(c)
		rawUserID := sess.Get("user_id")

		if rawUserID != nil {
			if userID, err := uuid.Parse(rawUserID.(string)); err == nil {
				if user, err := authSvc.GetByID(c.Request.Context(), userID.String()); err == nil {
					c.Set(ContextUserKey, user)
				}
			}
		} else {
			delete(c.Keys, "user_id")
		}
		c.Next()
	}
}

func (s *Server) AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, exists := c.Get(ContextUserKey); exists {
			c.Next()
		} else {
			c.Redirect(http.StatusSeeOther, "/auth/google")
		}
	}
}
