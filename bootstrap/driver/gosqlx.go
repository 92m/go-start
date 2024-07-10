package driver

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/92m/go-start/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	SqlxDB           *sqlx.DB
	sqlxErr          error
	sqldb_mysql_once sync.Once
)

// 创建全局sqlx数据库
func InitSqlx() {
	sqldb_mysql_once.Do(func() {
		if SqlxDB == nil {
			dbConfig := config.GetMySQLConfig()
			driverName := "mysql"
			dataSourceName := fmt.Sprintf(`%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&loc=Local`,
				dbConfig["DB_USERNAME"],
				dbConfig["DB_PASSWORD"],
				dbConfig["DB_HOST"],
				dbConfig["DB_PORT"],
				dbConfig["DB_SCHEMA"],
				dbConfig["DB_CHARSET"],
			)

			// fmt.Printf("driverName %s \n", driverName)
			// fmt.Printf("dataSourceName %s \n", dataSourceName)

			SqlxDB, sqlxErr = sqlx.Connect(driverName, dataSourceName)

			if sqlxErr != nil {
				log.Println("sqlx数据库连接失败。。。", sqlxErr)
				os.Exit(200)
			} else {
				log.Println("sqlx已连接现有数据库驱动 >>> ")
			}
		}
	})
}
