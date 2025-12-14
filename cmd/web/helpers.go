package main

import (
	"crypto/rand"
	"encoding/base64"
	"log"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func newLogger() *zap.SugaredLogger {
	// ---- Zap Logger ----
	cfg := zap.NewProductionConfig()
	zapLogger, err := cfg.Build()
	if err != nil {
		log.Fatal(err)
	}
	defer zapLogger.Sync()

	logger := zapLogger.Sugar()
	return logger
}

func (s *Server) render(c *gin.Context, status int, tmpl string, data gin.H) {
	if user, ok := c.Get(ContextUserKey); ok && user != nil {
		data["user"] = user
	}

	if flashes, ok := c.Get(ContextFlashKey); ok {
		data["flash"] = flashes
	}
	c.HTML(status, tmpl, data)
}

func generateState() string {
	b := make([]byte, 32)
	_, _ = rand.Read(b)
	return base64.RawURLEncoding.EncodeToString(b)
}
