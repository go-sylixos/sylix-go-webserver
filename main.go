package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(LoggerToFile())

	run(router)
	addr := fmt.Sprintf(":%d", 10000)
	Logger.Infof("Listening on http://%v", addr)
	Logger.Fatal(router.Run(addr))
}

//go:embed static/*
var staticFS embed.FS

func run(router *gin.Engine) { // http://localhost:10000/static/index.html
	static, _ := fs.Sub(staticFS, "static")
	router.StaticFS("/static", http.FS(static))

	// 主页路由
	router.GET("", func(c *gin.Context) {
		content, err := staticFS.ReadFile("static/index.html")
		if err != nil {
			c.String(http.StatusNotFound, "index.html not found")
			return
		}

		c.Data(http.StatusOK, "text/html; charset=utf-8", content)
	})

	router.NoRoute(gin.WrapH(http.FileServer(http.FS(static))))

	aclApis := router.Group("")
	AclRoute(aclApis)

	apiGroup := router.Group("/api")

	apiGroup.GET("/index", func(c *gin.Context) {
		c.String(http.StatusOK, "welcoming webserver!")
	})

	bizApis := apiGroup.Group("/v1")
	BizRoute(bizApis)
}
