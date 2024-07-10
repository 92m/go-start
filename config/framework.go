package config

import (
	"os"
	"runtime"
)

func GetFrameworkConfig() map[string]string {

	viper := CreateConfig("configuration")
	viper.WatchConfig()

	// main.go文件的绝对路径
	mainDirectory, _ := os.Getwd()
	mainDirectory = mainDirectory + "/"
	timezone := viper.GetString("framework.timezone")
	glVersion := viper.GetString("glVersion")
	logDebug := viper.GetString("confing.logDebug")

	if len(timezone) == 0 {
		timezone = "Asia/Shanghai"
	}
	if len(glVersion) == 0 {
		glVersion = "gl_version.default.0.0.0"
	}
	if len(logDebug) == 0 || (logDebug != "true" || logDebug != "false") {
		logDebug = "true"
	}
	conf := make(map[string]string)

	conf["timezone"] = timezone    // 时区
	conf["gl_version"] = glVersion // GinLaravel（或Ginvel）版本信息

	conf["go_version"] = runtime.Version() // go版本
	conf["go_root"] = runtime.GOROOT()

	conf["framework_path"] = mainDirectory            // 默认使用框架根目录
	conf["storage_path"] = mainDirectory + "storage/" // 文件存储文件夹

	conf["log_debug"] = logDebug

	return conf
}
