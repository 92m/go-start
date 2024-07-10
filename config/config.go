package config

import (
	"fmt"
	"log"
	"os"
	"path"
	"runtime"

	"github.com/spf13/viper" // 相对于$GOPATH/pkg/mod的路径
)

var (
	ProjectRootPath = path.Dir(getoncurrentPath()+"/../") + "/"
)

/*
 * 返回文件所在的目录
 * Caller 返回前期执行代码文件位置
 */

func getoncurrentPath() string {
	// Caller 参数 0表示当前本行代码在什么位置
	_, filename, _, _ := runtime.Caller(0)
	return path.Dir(filename)
}

/*
 * 利用viper库解析配置文件
 * Viper可以解析JSON、TOML、YAML、HCL、INI、ENV等格式的配置文件，甚至可以监听配置文件的变化(WatchConfig)，不需要重启程序就可以读到最新的值
 * 单元测试：ENVIRONMENT=prod go test -v  -timeout 30s -run ^TestConfig$ blog/util/test
 */

func CreateConfig(file string) *viper.Viper {
	config := viper.New()
	configPath := ProjectRootPath
	config.AddConfigPath(configPath) // 文件所在目录
	config.SetConfigName(file)       // 文件名
	config.SetConfigType("yaml")     // 文件类型
	configFile := configPath + file + ".yaml"

	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// panic(fmt.Errorf("找不到配置文件:%s", configFile)) //系统初始化阶段发生任何错误，直接结束进程
			log.Println(fmt.Errorf("找不到配置文件:%s", configFile))
			os.Exit(200)
		} else {
			// panic(fmt.Errorf("解析配置文件%s出错:%s", configFile, err))
			log.Println(fmt.Errorf("解析配置文件%s出错:%s", configFile, err))
			os.Exit(200)
		}
	}

	return config
}
