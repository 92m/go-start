package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/92m/go-start/util"
)

func TestProjectRootPath(t *testing.T) {
	fmt.Println(util.ProjectRootPath)
}

func TestConfig(t *testing.T) {
	dbViper := util.CreateConfig("mysql")
	dbViper.WatchConfig()

	// 读取配置的第一种方式
	if !dbViper.IsSet("blog.port") {
		t.Fail()
	}

	port := dbViper.GetInt("blog.port")
	fmt.Printf("blog.port 1: %d \n", port)

	// 10秒之内修改一下配置文件，看看viper能不能读取最新值
	time.Sleep(10 * time.Second)
	port = dbViper.GetInt("blog.port")
	fmt.Printf("blog.port 2: %d \n", port)

	// 读取配置的第二种方式
	logViper := util.CreateConfig("log")
	logViper.WatchConfig()

	type LogConfig struct {
		Level string `mapstructure:"level"` //Tag
		File  string `mapstructure:"file"`
	}

	var config LogConfig
	if err := logViper.Unmarshal(&config); err != nil {
		fmt.Println(err)
		t.Fail()
	} else {
		fmt.Printf("config.Level: %s \n", config.Level)
		fmt.Printf("config.File: %s \n", config.File)
	}
}

// 测试指令
// go test -v ./util/test/ -run=^TestProjectRootPath$ -count=1
// go test -v ./util/test/ -run=^TestConfig$ -count=1
