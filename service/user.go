package service

import (
	"hyper-manage/database"
	"hyper-manage/middleware"
	"hyper-manage/model"
	"hyper-manage/utils"
	"log"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserLogin struct {
	UserName string `form:"user_name"`
	Password string `form:"password"`
}

func (data *UserLogin) Login(c *gin.Context) {
	var user model.User

	if err := database.DB.Where("user_name = ?", data.UserName).First(&user).Error; err != nil {
		utils.Warning(c, "用户名或密码错误")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		utils.Warning(c, "用户名或密码错误")
		return
	}

	token, err := middleware.ReleaseToken(user)
	if err != nil {
		log.Print("Token generate error:", err)
		utils.Error(c, "Token分发失败")
		return
	}
	utils.Success(c, gin.H{"token": token}, "登录成功")
}

type UserRegister struct {
	UserName string `form:"user_name"`
	Password string `form:"password"`
	Super    bool   `form:"super"`
}

func (data *UserRegister) Register(c *gin.Context) {
	if len(data.UserName) == 0 || len(data.Password) == 0 {
		utils.Warning(c, "用户名&密码不应为空")
		return
	}
	// 判断用户名是否存在
	count := int64(0)
	database.DB.Model(&model.User{}).Where("user_name = ?", data.UserName).Count(&count)
	if count > 0 {
		utils.Warning(c, "用户已存在")
		return
	}
	// 创建用户
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Print("Password Hash error:", err)
		utils.Error(c, "密码加密失败")
		return
	}
	newUser := model.User{
		UserName: data.UserName,
		Password: string(hashPassword),
		Super:    data.Super,
	}
	database.DB.Create(&newUser)
	utils.Success(c, nil, "注册成功")
}
