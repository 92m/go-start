package test

import (
	"fmt"
	"testing"

	"github.com/92m/go-start/app/Common"
)

func TestGetServerInfo(t *testing.T) {

	if len(Common.ServerInfo["timezone"]) == 0 {
		t.Fail()
	}

	fmt.Printf("Common.ServerInfo[timezone]: %s", Common.ServerInfo["timezone"])
}

// 测试指令
// go test -v ./app/test -run=^TestGetServerInfo$ -count=1
