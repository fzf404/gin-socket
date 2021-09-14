package main

import (
	"fmt"
	"hyper-manage/test"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	assert := assert.New(t)
	router := Setup()

	w := test.Get("/ping", router)
	assert.Equal(200, w.Code)
	assert.Equal("pong", w.Body.String())

	formData := map[string]string{
		"user_name": "fzf",
		"password":  "12345678",
	}

	w = test.PostForm("/user/login", formData, router)
	assert.Equal(200, w.Code)
	fmt.Println(w.Body)
}
