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
	if isUser, err := db.CheckUserExists(s, c.Param("username")); isUser && err == nil {

		// Get user profile
		if user, err := db.GetUser(s, c.Param("username")); err == nil {
			c.HTML(
				// Set the HTTP status to 200 (OK)
				http.StatusOK,
				// Use the index.html template
				"results.html",
				// Pass the data that the page uses
				gin.H{
					"name":    user.Name,
					"payload": user,
				},
			)

		} else {
			// If the user is not found, abort with an error
			c.AbortWithError(http.StatusNotFound, err)
		}

	} else {
		// If an invalid username is specified in the URL, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
	}

}
