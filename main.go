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

// 初始化给定的 gin.Engine 路由器，设置各种路由和静态文件处理。
//
// 它设置了以下路由：
// - 从 "static" 目录在 "/static" URL 路径提供静态文件。
// - 一个根路由，从 "static" 目录提供 "index.html" 文件。
// - 一个处理 404 错误的回退路由，通过提供静态文件。
// - 一个 ACL（访问控制列表）路由组。
// - 一个带有 "/index" 端点的 API 路由组，返回欢迎消息。
// - 一个位于 "/api/v1" 下的业务 API 路由组。
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
