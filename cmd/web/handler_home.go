package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) Home(c *gin.Context) {
	cats, err := s.CategoryService.ListCategories(c.Request.Context())
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	s.render(c, 200, "home.tmpl", gin.H{
		"categories": cats,
	})
}
