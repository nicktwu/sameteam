package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nicktwu/sameteam/backend/middleware"
	"github.com/nicktwu/sameteam/backend/routes"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"

	"github.com/northbright/keygen"
)

func main() {

	key, err := keygen.GenSymmetricKey(256)
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	session, err := mgo.Dial("db:8080")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	r.Use(middleware.WithDB(session))
	r.POST("/register", routes.RegisterUser(key))
	r.Use(middleware.WithValidation(key))

	r.POST("/login", routes.GetLogin(key))
	r.GET("/ping", routes.Home)
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})
	r.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})

	r.Run()
}
