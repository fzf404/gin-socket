package service

import (
	"hyper-manage/database"
	"hyper-manage/model"
	"hyper-manage/utils"

	"github.com/gin-gonic/gin"
)

type ServiceInfo struct {
	ProjectName string `form:"project_name"`
	RepoUrl     string `form:"repo_url"`
	Branch      string `form:"branch"`
	InstallCmd  string `form:"install_cmd"`
	BuildCmd    string `form:"build_cmd"`
	OutputPath  string `form:"output_path"`
}

func (data ServiceInfo) New(c *gin.Context) {

	if len(data.ProjectName) == 0 || len(data.RepoUrl) == 0 || len(data.Branch) == 0 {
		utils.Warning(c, "数据不符合规范")
		return
	}

	if len(data.OutputPath) == 0 {
		data.OutputPath = "."
	}

	count := int64(0)
	database.DB.Model(&model.Service{}).Where("project_name = ?", data.ProjectName).Count(&count)
	if count > 0 {
		utils.Warning(c, "仓库名已存在")
		return
	}

	// 复制结构体
	newService := model.Service{}
	utils.StructAssign(&newService, &data)

	// 保存
	database.DB.Create(&newService)

	user, _ := c.Get("user")
	utils.Success(c, gin.H{"user_name": user.(*model.User).UserName, "url": "m5xcf9fg"}, "")
}

func (data ServiceInfo) Deploy(c *gin.Context) {
	utils.Success(c, gin.H{"url": "m5xcf9fg"}, "")
}
