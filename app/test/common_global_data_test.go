package test

import (
	"fmt"
	"testing"

	"github.com/92m/go-start/app/Common"
)

// 测试设置全局属性
func TestSetGlobalData(t *testing.T) {
	Common.SetGlobalData("test", "2021")
	if Common.GetGlobalData("test") != "2021" {
		t.Fail()
	}
	var test interface{} = Common.GetGlobalData("test")
	fmt.Printf("SetGlobalData:%s \n", test)
}

// 测试查询全局属性
func TestGetGlobalData(t *testing.T) {
	var test interface{} = Common.GetGlobalData("test")
	fmt.Println(test)
}

// 测试删除全局属性
func TestDelGlobalData(t *testing.T) {
	Common.DelGlobalData("test")
	var test interface{} = Common.GetGlobalData("test")

	if test != nil {
		t.Fail()
	}

	fmt.Printf("DelGlobalData: test \n")
}

// 测试指令
// go test -v ./app/test -run=^TestSetGlobalData$ -count=1
// go test -v ./app/test -run=^TestGetGlobalData$ -count=1
// go test -v ./app/test -run=^TestDelGlobalData$ -count=1
