package routes

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/nicktwu/sameteam/backend/db"
	"github.com/nicktwu/sameteam/backend/middleware"
	"github.com/nicktwu/sameteam/backend/models"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
)

func GetLogin(key []byte) func(c *gin.Context) {
	return func(c *gin.Context) {
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

		var user models.User

		if err := c.BindJSON(&user); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		exists, err := db.CheckUserExists(s, user.Username)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		if !exists {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		dbUser, err := db.GetUser(s, user.Username)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		if err := bcrypt.CompareHashAndPassword(dbUser.PasswordHash, bytes.NewBufferString(user.Password).Bytes()); err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token, err := middleware.GetToken(user.Username, key)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusAccepted, gin.H{"token": token, "user": dbUser})
	}
}

func RegisterUser(key []byte) func(c *gin.Context) {
	return func(c *gin.Context) {
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

		var user models.User

		if err := c.BindJSON(&user); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		log.Printf("User object is: %v", user)

		exists, err := db.CheckUserExists(s, user.Username)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		if exists {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		} else {
			if user.Username == "" {
				c.AbortWithStatus(http.StatusUnavailableForLegalReasons) // copyright Nick Wu
				return
			}

			if len(user.Password) < 8 {
				c.AbortWithStatus(http.StatusUpgradeRequired) // upgrade to longer password
				return
			}

			hash, err := bcrypt.GenerateFromPassword(bytes.NewBufferString(user.Password).Bytes(), bcrypt.DefaultCost)
			if err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
			user.PasswordHash = hash

			if err := db.InsertUser(s, user); err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
		}

		token, err := middleware.GetToken(user.Username, key)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusAccepted, gin.H{"token": token, "user": user})
	}
}
