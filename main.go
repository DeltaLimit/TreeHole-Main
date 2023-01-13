package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func init() {

}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"urIP": c.ClientIP(),
		})
	})
	err := router.Run()
	HandleError(err)
}

func HandleError(e error) {
	log.Fatalln(e)
	return
}
