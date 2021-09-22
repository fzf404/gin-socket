/*
 * @Author: fzf404
 * @Date: 2021-09-22 17:01:07
 * @LastEditTime: 2021-09-22 19:44:29
 * @Description: description
 */
package model

import "gorm.io/gorm"

type WebApp struct {
	gorm.Model
	ProjectName string
	RepoUrl     string
	Branch      string
	InstallCmd  string
	BuildCmd    string
	OutputPath  string
}
