package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nicktwu/sameteam/backend/db"
	"gopkg.in/mgo.v2"
	"net/http"
)

func ShowUserResults(c *gin.Context) {
	session, exists := c.Get("db")
	if !exists {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	s, valid := session.(*mgo.Session)
	if !valid {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	// Check if the user exists
	isUser, err := db.CheckUserExists(s, c.Param("username"))
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if !isUser {
		c.AbortWithStatus(http.StatusNotFound)
	}

	// Get user profile
	if user, err := db.GetUser(s, c.Param("username")); err == nil {
		c.JSON(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			// Pass the data that the page uses
			user,
		)
	} else {
		// If the user is not found, abort with an error
		c.AbortWithError(http.StatusNotFound, err)
	}

}
