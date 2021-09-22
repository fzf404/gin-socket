/*
 * @Author: fzf404
 * @Date: 2021-09-21 22:47:44
 * @LastEditTime: 2021-09-22 15:55:52
 * @Description: 测试
 */
package test

import (
	"bytes"
	"encoding/json"
	"gin-socket/service"
	"mime/multipart"
	"net/http/httptest"
	"net/url"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

/**
 * @description: map解析为字符串
 */
func ParseToStr(mp map[string]string) string {
	values := ""
	for key, val := range mp {
		values += "&" + key + "=" + val
	}
	temp := values[1:]
	values = "?" + temp
	return values
}

/**
 * @description: 发送Get请求
 */
func Get(url string, router *gin.Engine) *httptest.ResponseRecorder {
	// 构造请求
	req := httptest.NewRequest("GET", url, nil)
	// 初始化响应
	w := httptest.NewRecorder()
	// 调用接口
	router.ServeHTTP(w, req)
	return w
}

/**
 * @description: Post请求-form-data
 */
func PostForm(url string, formData map[string]string, router *gin.Engine) *httptest.ResponseRecorder {
	// 创建表单
	buf := new(bytes.Buffer)
	mw := multipart.NewWriter(buf)
	for key, value := range formData {
		mw.WriteField(key, value)
	}
	mw.Close()

	req := httptest.NewRequest("POST", url, buf)
	req.Header.Set("Content-Type", mw.FormDataContentType())
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}

/**
 * @description: Post请求-urlencoded
 */
func PostEncoded(url string, fromEncoded url.Values, router *gin.Engine) *httptest.ResponseRecorder {
	req := httptest.NewRequest("POST", url, strings.NewReader(fromEncoded.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}

/**
 * @description: Post请求-json
 */
func PostJson(url string, param map[string]interface{}, router *gin.Engine) *httptest.ResponseRecorder {
	// 转换为byte
	jsonByte, _ := json.Marshal(param)

	req := httptest.NewRequest("POST", url, bytes.NewReader(jsonByte))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}

func TestSocket() {
	for {
		service.SendClientSocket("home", "hi")
		time.Sleep(time.Second)
	}
}
