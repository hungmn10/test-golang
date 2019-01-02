
package main

import (
	"fmt"
	//"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Welcome to the playground!")
    router := gin.Default()
    router.GET("/user/:name", func(c *gin.Context) {
        name := c.Param("name")
        c.String(http.StatusOK, "Hello %s", name)
	}) 

	router.GET("/user", func(c *gin.Context) {
        id := c.Query("id")
        c.String(http.StatusOK, "Hello id%s", id)
	})



	router.POST("/post", func(c *gin.Context) {

        id := c.Query("id")
        page := c.DefaultQuery("page", "0")
        name := c.PostForm("name")
        message := c.PostForm("message")

        fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
    })
	router.Run(":8080")


 }