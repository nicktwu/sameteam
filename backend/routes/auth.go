package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nicktwu/sameteam/backend/db"
	"github.com/nicktwu/sameteam/backend/middleware"
	"github.com/nicktwu/sameteam/backend/models"
	"gopkg.in/mgo.v2"
	"net/http"
)

func GetLogin(key []byte) func(c *gin.Context) {
	return func(c *gin.Context) {
		session, exists := c.Get("db")
		if !exists {
			c.AbortWithStatus(http.StatusInternalServerError)
		}

		s, valid := session.(*mgo.Session)
		if !valid {
			c.AbortWithStatus(http.StatusInternalServerError)
		}

		var user models.User

		if err := c.BindJSON(&user); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
		}

		exists, err := db.CheckUserExists(s, user.Username)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
		}

		if !exists {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		token, err := middleware.GetToken(user.Username, key)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
		}

		c.JSON(http.StatusAccepted, gin.H{"token": token})
	}
}

func RegisterUser(key []byte) func(c *gin.Context) {
	return func(c *gin.Context) {
		session, exists := c.Get("db")
		if !exists {
			c.AbortWithStatus(http.StatusInternalServerError)
		}

		s, valid := session.(*mgo.Session)
		if !valid {
			c.AbortWithStatus(http.StatusInternalServerError)
		}

		var user models.User

		if err := c.BindJSON(&user); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
		}

		exists, err := db.CheckUserExists(s, user.Username)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
		}

		if exists {
			c.AbortWithStatus(http.StatusUnauthorized)
		} else {
			err := db.InsertUser(s, user)

			if user.Username == "" {
				c.AbortWithStatus(http.StatusUnavailableForLegalReasons) // copyright Nick Wu
			}

			passw := string(user.Password)
			if len(passw) < 8 {
				c.AbortWithStatus(http.StatusUpgradeRequired) // upgrade to longer password
			}
			if passw == "12345678" {
				fmt.Fprintln("fuk u")
			}

			if err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}

		token, err := middleware.GetToken(user.Username, key)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
		}

		c.JSON(http.StatusAccepted, gin.H{"token": token})
	}
}
