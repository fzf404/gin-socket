/*
 * @Author: fzf404
 * @Date: 2021-09-18 21:53:34
 * @LastEditTime: 2021-09-22 19:49:36
 * @Description: 应用服务
 */
package service

import (
	"gin-socket/database"
	"gin-socket/model"
	"gin-socket/utils"

	"github.com/gin-gonic/gin"
)

/**
* @description: Web应用
 */
type WebApp struct {
	ProjectName string `form:"project_name"`
	RepoUrl     string `form:"repo_url"`
	Branch      string `form:"branch"`
	InstallCmd  string `form:"install_cmd"`
	BuildCmd    string `form:"build_cmd"`
	OutputPath  string `form:"output_path"`
}

/**
 * @description: 新建Web应用
 */
func (data WebApp) New(c *gin.Context) {

	if len(data.ProjectName) == 0 || len(data.RepoUrl) == 0 || len(data.Branch) == 0 {
		utils.Warning(c, utils.CodeParamError, "数据不符合规范")
		return
	}

	if len(data.OutputPath) == 0 {
		data.OutputPath = "."
	}

	count := int64(0)
	database.DB.Model(&model.WebApp{}).Where("project_name = ?", data.ProjectName).Count(&count)
	if count > 0 {
		utils.Warning(c, utils.CodeParamError, "仓库名已存在")
		return
	}

	// 复制结构体
	newApp := model.WebApp{}
	utils.StructAssign(&newApp, &data)

	// 保存
	database.DB.Create(&newApp)

	utils.Success(c, gin.H{"app": data.ProjectName}, "新建成功")
}

/**
 * @description: 部署Web应用
 */
func (data WebApp) Deploy(c *gin.Context) {
	utils.Success(c, gin.H{}, "部署中...")
}
