package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

const (
	title      = "树洞系统"
	githubAddr = "https://github.com/DeltaLimit/Treehole-Main/"
)

func InitRouter(router *gin.Engine) {
	router.LoadHTMLGlob("static/html/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":       title,
			"github_addr": githubAddr,
			"time":        time.Now(),
		})
	})
}
