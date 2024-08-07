package routes

import (
	"log"

	"github.com/92m/go-start/routes/RouterGroup"
	"github.com/gin-gonic/gin"
)

// RouterRegister 灵活的注册路由文件
func RouterRegister(app *gin.Engine) {
	log.Println("运行自定义注册路由文件 >>> ")
	// 示例
	RouterGroup.ExampleApi(app) // 面向Api
	// 其他
	// RouterGroup.ApiGen1(app)
	// RouterGroup.ApiGen3(app)
}
