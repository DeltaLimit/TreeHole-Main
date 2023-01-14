package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func InitRouter(router *gin.Engine) {
	router.LoadHTMLGlob("static/html/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":       "树洞系统",
			"github_addr": "https://github.com/DeltaLimit/TreeHole-Main/",
			"time":        time.Now(),
		})
	})
}
