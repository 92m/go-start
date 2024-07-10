package bootstrap

import (
	"log"
	"os"

	"github.com/92m/go-start/app/Http/Middlewares"
	"github.com/92m/go-start/config"
	"github.com/gin-gonic/gin"
)

// App 配置并启动http服务
// 项目访问入口

func App(HttpServer *gin.Engine) {
	serverConfig := config.GetServerConfig()

	// Gin服务
	HttpServer = gin.Default()

	// 捕捉接口运行耗时（必须排第一）
	HttpServer.Use(Middlewares.StatLatency)

	// 设置全局ctx参数（必须排第二）
	HttpServer.Use(Middlewares.AppData)

	// 实际访问网址和端口
	_host := "127.0.0.1:" + serverConfig["PORT"]              // 测试访问IP
	host := serverConfig["HOST"] + ":" + serverConfig["PORT"] // Docker访问IP

	// glVersion := frameworkConfig["gl_version"]

	// 终端提示
	log.Println(
		//"\n App Success! \n\n " +
		" \n " +
			"访问地址示例：" + host + " >>> \n " +
			// "gl_version：" + glVersion + " \n " +
			"1) 默认接口（JSON输出）：http://" + _host + " \n " +
			"2) 测试接口（JSON输出）：http://" + _host + "/api?name=gl&id=2021 \n " +
			"3) 静态文件输出（文件）：http://" + _host + "/favicon.ico \n " +
			"4) 查看WebSocket连接：ws://" + _host + "/api/example/socket/ping1 \n " +
			"")

	err := HttpServer.Run(host)

	if err != nil {
		log.Println("http服务遇到错误，运行中断，error：", err.Error())
		log.Println("提示：注意端口被占时应该首先更改对外暴露的端口，而不是微服务的端口。")
		os.Exit(200)
	}
}
