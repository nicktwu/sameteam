package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nicktwu/sameteam/backend/middleware"
	"github.com/nicktwu/sameteam/backend/routes"
	"github.com/northbright/keygen"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
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

	r.Use(cors.Default())
	r.Use(middleware.WithDB(session))
	r.POST("/register", routes.RegisterUser(key))
	r.POST("/login", routes.GetLogin(key))

	api := r.Group("/api")
	api.Use(middleware.WithValidation(key))

	api.GET("/ping", routes.Home)
	api.GET("/user/:username", routes.ShowUserResults)
	api.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})

	r.Run()
}
