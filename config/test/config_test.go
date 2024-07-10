package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/92m/go-start/config"
)

func TestProjectRootPath(t *testing.T) {
	fmt.Println(config.ProjectRootPath)
}

func TestConfig(t *testing.T) {
	dbViper := config.CreateConfig("configuration")
	dbViper.WatchConfig()

	// 读取配置的第一种方式
	if !dbViper.IsSet("mysql.port") {
		t.Fail()
	}

	port := dbViper.GetInt("mysql.port")
	fmt.Printf("mysql.port 1: %d \n", port)

	// 10秒之内修改一下配置文件，看看viper能不能读取最新值
	time.Sleep(10 * time.Second)
	port = dbViper.GetInt("mysql.port")
	fmt.Printf("mysql.port 2: %d \n", port)
}

// 测试指令
// go test -v ./config/test/ -run=^TestConfig$ -count=1
