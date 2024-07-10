package config

func GetServerConfig() map[string]string {

	viper := CreateConfig("configuration")
	viper.WatchConfig()
	nodeName := "config.server"

	host := viper.GetString(nodeName + ".host")
	port := viper.GetString(nodeName + ".port")
	env := viper.GetString(nodeName + ".env")

	// 默认使用本地数据库服务
	// docker中运行请使用：0.0.0.0，本地测试请使用：127.0.0.1
	if host == "localhost" || len(host) == 0 {
		host = "0.0.0.0"
	}
	// 使用默认端口
	if len(port) == 0 {
		port = "3306"
	}

	if len(env) == 0 {
		env = "release"
	}

	conf := make(map[string]string)

	conf["HOST"] = host // 监听地址，部署在docker中请使用：0.0.0.0。建议不要用127.0.0.1或localhost
	conf["PORT"] = port // 监听端口
	conf["ENV"] = env   // 环境模式 release/debug/test

	return conf
}
