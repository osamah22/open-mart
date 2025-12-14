package main

import (
	"fmt"
	"net/http"

	"cloud.google.com/go/auth/credentials/idtoken"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap/zapcore"
	"golang.org/x/oauth2"
)

func (s *Server) GoogleLogin(c *gin.Context) {
	state := generateState()

	// ✅ THIS IS CRITICAL
	c.SetSameSite(http.SameSiteLaxMode)

	c.SetCookie(
		"oauth_state",
		state,
		300,
		"/",
		"",
		false, // Secure = false for localhost
		true,
	)

	url := s.OAuthConfig.AuthCodeURL(
		state,
		oauth2.AccessTypeOnline,
		oauth2.SetAuthURLParam("prompt", "select_account"),
		oauth2.SetAuthURLParam("prompt", "consent"),
	)

	c.Redirect(http.StatusFound, url)
}

func (s *Server) GoogleCallback(c *gin.Context) {
	ctx := c.Request.Context()

	// ---- CSRF state check ----
	state := c.Query("state")
	cookieState, err := c.Cookie("oauth_state")
	if err != nil || state != cookieState {
		s.Logger.Logf(zapcore.DebugLevel, "Unauthorized access: invalid state")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	code := c.Query("code")
	if code == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// ---- exchange code ----
	token, err := s.OAuthConfig.Exchange(ctx, code)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// ---- validate ID token ----
	validator, err := idtoken.NewValidator(nil)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	payload, err := validator.Validate(ctx, rawIDToken, s.OAuthConfig.ClientID)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// ---- extract claims safely ----
	googleID, _ := payload.Claims["sub"].(string)
	email, _ := payload.Claims["email"].(string)
	avatarURL, _ := payload.Claims["picture"].(string)
	fmt.Printf("google id %s", googleID)
	fmt.Printf("email %s", email)

	if googleID == "" || email == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// ---- get or create user ----
	user, err := s.authService.GetOrCreateUser(
		ctx,
		googleID,
		email,
		avatarURL,
	)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// ---- create session ----
	sess := sessions.Default(c)
	sess.Set("user_id", user.ID.String())
	sess.Set("avatar_url", avatarURL)

	if err := sess.Save(); err != nil {
		s.Logger.Error("session save failed", err)
	}

	fmt.Println("SESSION SAVED user_id =", sess.Get("user_id"))

	c.Redirect(http.StatusSeeOther, "/")
}

func (s *Server) Logout(c *gin.Context) {
	sess := sessions.Default(c)

	// 1. Clear session data
	sess.Clear()
	_ = sess.Save()
	sess.Set(ContextFlashKey, "info: Logged out successfully")

	// 3. Redirect
	c.Redirect(http.StatusSeeOther, "/")
}
