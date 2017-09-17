package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nicktwu/sameteam/backend/routes"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
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
