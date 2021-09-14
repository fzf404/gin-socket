package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func Success(c *gin.Context, data gin.H, msg string) {
	c.JSON(http.StatusOK, Response{
		Code: 200,
		Data: data,
		Msg:  msg,
	})
}

func Warning(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, Response{
		Code: 400,
		Data: nil,
		Msg:  msg,
	})
}

func Error(c *gin.Context, msg string) {
	c.JSON(http.StatusInternalServerError, Response{
		Code: 500,
		Data: nil,
		Msg:  msg,
	})
}
