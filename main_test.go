/*
 * @Author: fzf404
 * @Date: 2021-09-22 16:48:03
 * @LastEditTime: 2021-09-22 19:59:11
 * @Description: description
 */
package main

import (
	"gin-socket/test"
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
 * @description: 测试
 */
func TestRoute(t *testing.T) {
	assert := assert.New(t)
	router := Setup()

	w := test.Get("/ping", router)
	assert.Equal(200, w.Code)
	assert.Equal("pong", w.Body.String())

	formData := map[string]string{
		"user_name": "admin",
		"password":  "123456",
	}

	w = test.PostForm("/user/login", formData, router)
	assert.Equal(200, w.Code)
}
