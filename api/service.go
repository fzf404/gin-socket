/*
 * @Author: fzf404
 * @Date: 2021-09-22 14:08:27
 * @LastEditTime: 2021-09-22 19:50:02
 * @Description: Service接口
 */
package api

import (
	"gin-socket/service"
	"gin-socket/utils"

	"github.com/gin-gonic/gin"
)

func NewWebApp(c *gin.Context) {
	var data service.WebApp
	if err := c.ShouldBind(&data); err != nil {
		utils.Warning(c, utils.CodeParamError, "数据不符合规范")
		return
	}
	data.New(c)
}

func Socket(c *gin.Context) {
	name := c.Query("name")
	if len(name) == 0 {
		utils.Warning(c, utils.CodeParamError, "数据不符合规范")
		c.Abort()
		return
	}
	service.SocketServer(c.Writer, c.Request, name)
}
