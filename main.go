package main

import (
	router2 "TreeHole-Main/router"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	router2.InitRouter(router)
	err := router.Run()
	HandleError(err)
}

func HandleError(e error) {
	log.Fatalln(e)
	return
}
