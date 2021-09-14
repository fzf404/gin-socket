package main

import (
	"hyper-manage/database"
	"hyper-manage/router"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	database.InitSQLite()
	return router.NewRouter()
}

func main() {
	r := Setup()
	r.Run(":8080")
}
