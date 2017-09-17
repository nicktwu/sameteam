package routes

import "github.com/gin-gonic/gin"
import "gopkg.in/mgo.v2"

// Home
func Home(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// Form Submission
func Form(c *gin.Context) {
	c.
}

// Results
