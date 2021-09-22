/*
 * @Author: fzf404
 * @Date: 2021-09-14 14:45:06
 * @LastEditTime: 2021-09-22 16:50:36
 * @Description: Ping接口
 */
package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
