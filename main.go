package main

import (
	"TreeHole-Main/router"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	router.InitRouter(r)
	err := r.Run()
	HandleError(err)
}

func HandleError(e error) {
	log.Fatalln(e)
	return
}
