package ControllerGosqlx

import "github.com/gin-gonic/gin"

func ListUser(ctx *gin.Context) {
	// var state int
	// var msg string

	type ListUserKeys struct { // 结果集，参数名需大写
		id   int
		name int
	}

	var back = map[string]interface{}{
		"state":   1,
		"msg":     "接口请求成功，进入Gosqlx示例",
		"content": 1,
	}
	ctx.JSONP(200, back)
}
