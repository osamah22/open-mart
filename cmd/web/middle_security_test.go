package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSecureHeaders(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.New()
	r.Use(secureHeaders())
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t,
		"default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com 'self'",
		resp.Header.Get("Content-Security-Policy"),
	)
	assert.Equal(t, "origin-when-cross-origin", resp.Header.Get("Referrer-Policy"))
	assert.Equal(t, "nosniff", resp.Header.Get("X-Content-Type-Options"))
	assert.Equal(t, "deny", resp.Header.Get("X-Frame-Options"))
	assert.Equal(t, "0", resp.Header.Get("X-XSS-Protection"))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
