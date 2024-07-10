package test

import (
	"log"
	"testing"

	"github.com/92m/go-start/app/Common"
)

// 测试
func TestGetUrlParam(t *testing.T) {
	id := Common.GetUrlParam("baidu.com?date=2021&id=1949&name=牛逼", "id")
	if id != "1949" {
		t.Error("GetUrlParam()函数运行错误，test-id=", id)
	} else {
		log.Println("GetUrlParam()函数运行通过，test-id=", id)
		log.Println(Common.ServerInfo)
	}
}

// 测试指令
// go test -v ./app/test -run=^TestGetUrlParam$ -count=1
