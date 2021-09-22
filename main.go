/*
 * @Author: fzf404
 * @Date: 2021-09-22 15:49:29
 * @LastEditTime: 2021-09-22 16:22:13
 * @Description: 主函数
 */
package main

import (
	"gin-socket/database"
	"gin-socket/router"
	"gin-socket/service"

	"github.com/gin-gonic/gin"
)

/**
 * @description: 初始化数据库和路由
 */
func Setup() *gin.Engine {
	database.InitSQLite()
	service.SocketInit()
	return router.NewRouter()
}

func main() {
	r := Setup()
	r.Run(":8080")
}
