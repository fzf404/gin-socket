package api

import (
	"hyper-manage/service"
	"hyper-manage/utils"

	"github.com/gin-gonic/gin"
)

func NewService(c *gin.Context) {
	var data service.ServiceInfo
	if err := c.ShouldBind(&data); err != nil {
		utils.Warning(c, "数据不符合规范")
		return
	}
	data.New(c)
}
