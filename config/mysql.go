package config

import (
	"log"
	"os"
)

func GetMySQLConfig() map[string]string {

	dbViper := CreateConfig("configuration")
	dbViper.WatchConfig()
	dbName := "config.mysql"

	host := dbViper.GetString(dbName + ".host")
	schema := dbViper.GetString(dbName + ".schema")
	port := dbViper.GetString(dbName + ".port")
	username := dbViper.GetString(dbName + ".user")
	password := dbViper.GetString(dbName + ".pass")
	charset := dbViper.GetString(dbName + ".charset")

	// 默认使用本地数据库服务
	if len(host) == 0 {
		host = "127.0.0.1"
	}

	// 默认使用测试数据库
	if len(schema) == 0 {
		schema = "test"
	}

	// 使用默认端口
	if len(port) == 0 {
		port = "3306"
	}

	// 判断登陆信息
	if len(username) == 0 || len(password) == 0 {
		log.Println("mysql账户名或密码不能为空。。。")
		os.Exit(200)
	}

	if len(charset) == 0 {
		charset = "utf8mb4"
	}

	conf := make(map[string]string)

	conf["DB_HOST"] = host // 127.0.0.1
	conf["DB_PORT"] = port
	conf["DB_SCHEMA"] = schema
	conf["DB_USERNAME"] = username
	conf["DB_PASSWORD"] = password
	conf["DB_CHARSET"] = charset // "utf8mb4"
	conf["DB_TIMEOUT"] = "12s"

	conf["DB_MAX_OPEN_CONNS"] = "20"       // 连接池最大连接数
	conf["DB_MAX_IDLE_CONNS"] = "10"       // 连接池最大空闲数
	conf["DB_MAX_LIFETIME_CONNS"] = "7200" // 连接池链接最长生命周期

	return conf
}
