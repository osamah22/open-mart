package main

import (
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/osamah22/open-mart/internal/auth"
)

const ContextUserKey = "user"

func AuthContext(authSvc auth.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		sess := sessions.Default(c)
		rawUserID := sess.Get("user_id")

		fmt.Printf("auth middleware\n")
		fmt.Printf("user_id: %#v\n", rawUserID)
		if rawUserID != nil {
			if userID, err := uuid.Parse(rawUserID.(string)); err == nil {
				if user, err := authSvc.GetByID(c.Request.Context(), userID.String()); err == nil {
					fmt.Printf("username: %s", user.Username)
					c.Set(ContextUserKey, user)
				}
			}
		} else {
			delete(c.Keys, "user_id")
		}
		c.Next()
	}
}

func AuthRequired(authSvc auth.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, exists := c.Get(ContextUserKey); exists {
			c.Next()
		} else {
			c.Redirect(422, "/auth/google")
		}
	}
}
