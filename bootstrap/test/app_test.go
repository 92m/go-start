package test

import (
	"fmt"
	"testing"

	"github.com/92m/go-start/bootstrap"
	"github.com/gin-gonic/gin"
)

var HttpServer *gin.Engine

func TestRunApp(t *testing.T) {

	fmt.Printf("Gin* main.go】启动main.go... \n")
	// 启动服务
	bootstrap.App(HttpServer)
}

// 测试指令
// go test -v ./bootstrap/test/ -run=^TestRunApp$ -count=1
