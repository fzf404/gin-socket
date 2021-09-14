package middleware

import (
	"hyper-manage/database"
	"hyper-manage/model"
	"hyper-manage/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取 Token
		tokenString := c.GetHeader("Authorization")

		// 判断Bearer
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			utils.Warning(c, "权限不足")
			c.Abort()
			return
		}

		// 判断Token
		tokenString = tokenString[7:]
		token, claims, err := ParseToken(tokenString)

		if err != nil || !token.Valid {
			utils.Warning(c, "权限不足")
			c.Abort()
			return
		}

		// 通过验证, 获取claims中的UserID
		userID := claims.UserID

		var user model.User
		database.DB.First(&user, userID)

		// 验证用户是否存在
		if user.ID == 0 {
			utils.Warning(c, "权限不足")
			c.Abort()
			return
		}

		// user信息写入上下文
		c.Set("user", &user)
		c.Next()
	}
}
