package test

import (
	"go-start/util"
	"testing"
)

func TestLogger(t *testing.T) {
	util.InitLog("log")
	util.LogRus.Debug("this is debug log")
	util.LogRus.Info("this is info log")
	util.LogRus.Warn("this is warn log")
	util.LogRus.Error("this is error log")
	// util.LogRus.Panic("this is panic log") //写完日志之后再调panic
}

// 测试指令
// go test -v ./util/test/ -run=^TestLogger$ -count=1
