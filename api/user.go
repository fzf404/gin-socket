package api

import (
	"fmt"
	"hyper-manage/service"
	"hyper-manage/utils"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var data service.UserLogin
	fmt.Print(c.Copy().Request.Body)
	if err := c.ShouldBind(&data); err != nil {
		utils.Warning(c, "用户名或密码错误")
		return
	}
	data.Login(c)
}

func Register(c *gin.Context) {
	var data service.UserRegister
	if err := c.ShouldBind(&data); err != nil {
		utils.Warning(c, "数据不符合规范")
		return
	}
	data.Register(c)
}
