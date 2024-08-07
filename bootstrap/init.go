package bootstrap

import (
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/92m/go-start/bootstrap/driver"
)

// go环境自动检测，不符合要求则会自动退出Ginvel服务
func init() {
	// 基础环境自检
	log.Println("【Gin* init()】系统基础环境自检...")
	var _goVersion string = runtime.Version()
	var minVersion string = "go1.15.0"
	var goVersion string = strings.Replace(_goVersion, "go", "", -1)
	var goArray []string = strings.Split(goVersion, ".")
	var go1 int
	var go2 int
	var goYes int

	for _, value := range goArray {
		theValue, _ := strconv.ParseInt(value, 10, 64)
		if go1 == 0 {
			go1 = int(theValue)
		} else {
			if go2 == 0 {
				go2 = int(theValue)
				if go1 == 1 { // min=1.15
					goYes = goYes + 1
					if go2 >= 15 {
						goYes = goYes + 1
					}
				} else if go1 >= 2 {
					goYes = goYes + 2
				}
			} else {
				break
			}
		}
	}
	if goYes == 2 {
		log.Println("Go版本符合要求 >>> 当前Go版本：", _goVersion, " 最小要求版本："+minVersion)
	} else {
		log.Println("Go版本太低，框架所依赖的特性将不可直接使用！服务自动中断！", "当前Go版本："+_goVersion, " 最小要求版本："+minVersion, "运行中断")
		os.Exit(200)
	}

	// 必选初始化
	driver.InitSqlx()
}
