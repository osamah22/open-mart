package main

import (
	"encoding/json"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var ContextFlashKey = "flash"

func flashMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		sess := sessions.Default(c)

		raw, _ := sess.Get(FlashKey).(string)
		if raw != "" {
			var flashes []Flash
			if err := json.Unmarshal([]byte(raw), &flashes); err == nil && len(flashes) > 0 {
				c.Set("flashes", flashes)
			}
			sess.Delete(FlashKey)
			_ = sess.Save()
		}

		c.Next()
	}
}
