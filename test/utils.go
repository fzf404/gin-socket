package test

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http/httptest"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

// map => query
func ParseToStr(mp map[string]string) string {
	values := ""
	for key, val := range mp {
		values += "&" + key + "=" + val
	}
	temp := values[1:]
	values = "?" + temp
	return values
}

func Get(url string, router *gin.Engine) *httptest.ResponseRecorder {
	// 构造请求
	req := httptest.NewRequest("GET", url, nil)
	// 初始化响应
	w := httptest.NewRecorder()
	// 调用接口
	router.ServeHTTP(w, req)
	return w
}

// form-data
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

// form-data + file
func PostFormFile(url string, formData map[string]io.Reader, router *gin.Engine) *httptest.ResponseRecorder {
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

// form-urlencoded
func PostEncoded(url string, fromEncoded url.Values, router *gin.Engine) *httptest.ResponseRecorder {
	req := httptest.NewRequest("POST", url, strings.NewReader(fromEncoded.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}

// json
func PostJson(url string, param map[string]interface{}, router *gin.Engine) *httptest.ResponseRecorder {
	// 转换为byte
	jsonByte, _ := json.Marshal(param)

	req := httptest.NewRequest("POST", url, bytes.NewReader(jsonByte))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}
