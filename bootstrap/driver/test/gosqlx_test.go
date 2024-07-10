package test

import (
	"fmt"
	"testing"

	"github.com/92m/go-start/bootstrap/driver"
)

/*
 * 标题：创建数据库blog库连接池
 * 指令：go test -v ./bootstrap/driver/test -run=^TestInitSqlx$ -count=1
 */

type User struct {
	ID   int64  `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

func TestInitSqlx(t *testing.T) {
	driver.InitSqlx()

	if driver.SqlxDB == nil {
		fmt.Print("db connection is nil \n")
		t.Fail()
	}

	sqlStr := "SELECT id, name FROM user LIMIT 10"

	var users []User

	err := driver.SqlxDB.Select(&users, sqlStr)

	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		t.Fail()
	}

	// // fmt.Printf("id:%d name:%s \n", u.ID, u.Name)
	fmt.Printf("users:%#v\n", users)
}
