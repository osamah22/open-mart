package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"log"

	"github.com/gin-contrib/sessions"
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
	if data == nil {
		data = gin.H{}
	}
	if flashes, ok := c.Get("flashes"); ok {
		data["flashes"] = flashes
	}

	if user, ok := c.Get(ContextUserKey); ok {
		data["user"] = user
	}

	c.HTML(status, tmpl, data)
}

func generateState() string {
	b := make([]byte, 32)
	_, _ = rand.Read(b)
	return base64.RawURLEncoding.EncodeToString(b)
}

func (s *Server) addFlash(c *gin.Context, typ, msg string) {
	sess := sessions.Default(c)

	// read existing
	var flashes []Flash
	if raw, ok := sess.Get(FlashKey).(string); ok && raw != "" {
		_ = json.Unmarshal([]byte(raw), &flashes)
	}

	// append
	flashes = append(flashes, Flash{Type: typ, Message: msg})

	// write back
	b, _ := json.Marshal(flashes)
	sess.Set(FlashKey, string(b))
	_ = sess.Save()
}

func (s *Server) ErrorMessage(c *gin.Context, msg string) {
	s.addFlash(c, "error", msg)
}

func (s *Server) InfoMessage(c *gin.Context, msg string) {
	s.addFlash(c, "info", msg)
}

func (s *Server) SuccessMessage(c *gin.Context, msg string) {
	s.addFlash(c, "success", msg)
}
