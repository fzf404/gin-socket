package router

import (
	"hyper-manage/api"
	"hyper-manage/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	route := r.Group("/")
	{
		route.GET("ping", api.Ping)
		route.POST("user/login", api.Login)
		route.POST("user/register", api.Register)
		route.Use(middleware.Auth())
		{
			route.POST("service/new", api.NewService)
		}
	}
	return r
}
