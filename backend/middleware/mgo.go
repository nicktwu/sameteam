package middleware

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

func WithDB(s *mgo.Session) func(c *gin.Context) {
	return func(c *gin.Context) {
		session := s.Clone()
		defer session.Close()
		c.Set("db", session)
		c.Next()
	}
}
