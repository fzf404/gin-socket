# Gin + websocket

> 使用gin为基础的websocket开发模板

- 鉴权：golang-jwt/jwt

- websocket：gorilla/websocket

- 测试：stretchr/testify

- ORM：gorm
- 数据库：sqlite

## 使用

```bash
go run main.go	# 运行
go test ./...	# 测试
```

## 模板接口

```bash
# /Ping
$ curl http://127.0.0.1:8080/ping
pong⏎  

# /user/login
$ curl -X POST -F 'user_name=admin' -F "password=123456" http://127.0.0.1:8080/user/login
{"code":200,"data":{"token":"eyJ...rPw"},"msg":"登录成功"}⏎

# /user/register
$ curl -X POST -H "Authorization:Bearer eyJ..Token..rPw" -F "user_name=admin" -F "password=12345678" http://127.0.0.1:8080/user/register
{"code":402,"data":null,"msg":"用户已存在"}

# /websocket
$ npm i -g wscat	# websocket测试工具
$ wscat -H "Authorization:Bearer eyJ..<Login的Token>..rPw" -c "ws://127.0.0.1:8080/websocket?name=log"
Connected (press CTRL+C to quit)
> 
# 此时请求Login
curl -X POST -F 'user_name=admin' -F "password=123456" http://127.0.0.1:8080/user/login
# websocket将获得值
< admin: 已登入
```

## 目录

```bash
# 详细注释查看文件头
.
├── api	# API接口
│   ├── main.go
│   ├── service.go
│   └── user.go
├── config	# 配置文件
├── database	# 数据库
│   ├── database.db	# 数据库文件
│   └── sqlite.go
├── middleware	# 中间件
│   ├── auth.go
│   ├── cros.go
│   └── jwt.go
├── model	# 数据库模型
│   ├── jwt.go
│   ├── service.go
│   └── user.go
├── router	# 路由
│   └── router.go
├── service	# API服务
│   ├── app.go
│   ├── user.go
│   └── websocket.go
├── test	# 测试
│   └── utils.go
├── utils	# 小工具
│   ├── response.go
│   └── utils.go
├── go.mod
├── go.sum
├── main.go	# 主函数
└── main_test.go	# 测试
```

