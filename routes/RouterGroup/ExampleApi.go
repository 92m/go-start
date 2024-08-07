package RouterGroup

import (
	"github.com/92m/go-start/app/Http/Controllers/Example"
	"github.com/92m/go-start/app/Http/Controllers/Example/ControllerGosqlx"
	"github.com/92m/go-start/app/Http/Middlewares"
	"github.com/gin-gonic/gin"
)

/*
路由访问原则：宽进严出，所以都用Any，在拦截器里面拦截（VerifyXXX.go）具体请求事件。
路由周期：请求路由名——>header过滤——>拦截请求频率——>校验请求方法和Token参数——>运行目标函数——>程序达到终点，关闭此次请求。
路由写法 ：Any(路由名（必选）, header参数（可选）, 访问频率限制（可选）, 拦截器参数验证（可选）, 目标函数handler（必选）)
路由命名原则：推荐使用4层路由。第1层：api类还是web类；第2层：接口版本名；第3层：不同拦截器下的不同空间命名；第4层：目标函数handler。
*/

// ExampleApi 面向模版。访问：你的域名/api/空间命名/具体方法
func ExampleApi(route *gin.Engine) {
	api := route.Group("/api/", Middlewares.HttpCorsApi)
	{
		example := api.Group("/example/")
		{
			gosqlx := example.Group("/gosqlx/", Middlewares.HttpLimiter(2), Example.VerifyExample)
			gosqlx.Any("list_user", ControllerGosqlx.ListUser)
		}
	}
}
