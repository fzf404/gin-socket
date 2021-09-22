/*
 * @Author: fzf404
 * @Date: 2021-09-18 21:57:17
 * @LastEditTime: 2021-09-21 22:48:09
 * @Description: 响应
 */
package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
 * @description: 响应结构体
 */
type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	CodeSuccess = 200 // 成功

	CodeNoLogin      = 401 // 未登录
	CodeParamError   = 402 // 参数错误
	CodeNoAuth       = 403 // 无权限
	CodeNotFind      = 404 // 找不到
	CodeTokenExpired = 405 // Token过期

	CodeTokenErr = 501 // Token分发失败
	CodeHashErr  = 502 // Token分发失败
)

/**
 * @description: 成功响应
 */
func Success(c *gin.Context, data gin.H, msg string) {
	c.JSON(http.StatusOK, Response{
		Code: CodeSuccess,
		Data: data,
		Msg:  msg,
	})
}

/**
 * @description: 提示响应
 */
func Warning(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusBadRequest, Response{
		Code: code,
		Data: nil,
		Msg:  msg,
	})
}

/**
 * @description: 错误响应
 */
func Error(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusInternalServerError, Response{
		Code: code,
		Data: nil,
		Msg:  msg,
	})
}
