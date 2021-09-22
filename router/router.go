/*
 * @Author: fzf404
 * @Date: 2021-09-22 14:11:02
 * @LastEditTime: 2021-09-22 19:57:44
 * @Description: 路由
 */
package router

import (
	"gin-socket/api"
	"gin-socket/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Cors())

	route := r.Group("/")
	{
		route.GET("ping", api.Ping)
		route.POST("user/login", api.Login)
		route.POST("user/register", api.Register)

		route.Use(middleware.Auth())
		{
			route.POST("app/new", api.NewWebApp)
			route.GET("websocket", api.Socket)
		}
	}
	return r
}
