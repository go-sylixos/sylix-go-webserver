package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// BizRoute 注册业务路由到指定的路由组。
// 参数:
//   - group: gin.RouterGroup 路由组对象。
func BizRoute(group *gin.RouterGroup) {
	group.GET("/app/list", list)
}

// list 处理分页请求并返回指定页的数据。
//
// 参数:
//   - c: gin.Context 上下文对象，用于获取请求参数和返回响应。
//
// 功能:
//   - 从请求中获取 pageNum 参数，表示请求的页码。
//   - 根据 pageNum 计算数据的起始和结束索引。
//   - 从 listData 中提取对应页的数据。
//   - 记录返回的数据条数。
//   - 返回包含分页数据的 JSON 响应。
//
// 返回:
//   - JSON 格式的响应，包含分页数据、页码、每页大小、总记录数、状态码和消息。
func list(c *gin.Context) {
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))

	start := (pageNum - 1) * 10
	end := start + 10
	if end >= len(listData) {
		end = len(listData)
	}

	data := listData[start:end]
	Logger.Debugf("%d records return", len(data))

	c.JSON(200, gin.H{
		"data": ListResponse{
			List:     data,
			PageNum:  pageNum,
			PageSize: 10,
			Total:    len(listData),
		},
		"status":  http.StatusOK,
		"message": "success",
	})
}

// AclRoute 设置ACL相关的路由组。
// 它定义了两个GET请求的路由：
// 1. /login - 返回登录成功的JSON响应。
// 2. /logout - 返回登出成功的JSON响应。
func AclRoute(group *gin.RouterGroup) {
	group.GET("/login", func(c *gin.Context) {
		if c.Query("username") == "" || c.Query("password") == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"msg":  "username or password is empty",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "success",
		})
	})

	group.GET("/logout", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "success",
		})
	})
}
