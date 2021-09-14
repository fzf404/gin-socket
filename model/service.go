package model

import "gorm.io/gorm"

type Service struct {
	gorm.Model
	ProjectName string
	RepoUrl     string
	Branch      string
	InstallCmd  string
	BuildCmd    string
	OutputPath  string
}
