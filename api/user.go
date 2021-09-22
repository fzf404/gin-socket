/*
 * @Author: fzf404
 * @Date: 2021-09-22 14:10:28
 * @LastEditTime: 2021-09-22 16:50:51
 * @Description: User接口
 */
package api

import (
	"gin-socket/service"
	"gin-socket/utils"

	"github.com/gin-gonic/gin"
)

/**
 * @description: 用户登录
 * @param {*gin.Context} c
 */
func Login(c *gin.Context) {
	var data service.UserLogin
	if err := c.ShouldBind(&data); err != nil {
		utils.Warning(c, utils.CodeParamError, "用户名或密码错误")
		return
	}
	data.Login(c)
}

/**
 * @description: 用户注册
 * @param {*gin.Context} c
 */
func Register(c *gin.Context) {
	var data service.UserRegister
	if err := c.ShouldBind(&data); err != nil {
		utils.Warning(c, utils.CodeParamError, "数据不符合规范")
		return
	}
	data.Register(c)
}
