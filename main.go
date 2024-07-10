package main

import (
	"log"

	"github.com/92m/go-start/bootstrap"
	"github.com/gin-gonic/gin"
)

var HttpServer *gin.Engine

func main() {
	log.Println("【Gin* main.go】启动main.go...")

	// 启动服务
	bootstrap.App(HttpServer)
}
