package config

func GetViewConfig() map[string]string {

	viper := CreateConfig("configuration")
	viper.WatchConfig()
	nodeName := "config.view"

	pattern := viper.GetString(nodeName + ".view_pattern")
	static := viper.GetString(nodeName + ".view_static")

	// 默认值
	if len(pattern) == 0 {
		pattern = "views/html/**/**/*"
	}
	if len(static) == 0 {
		static = "views/static/"
	}

	conf := make(map[string]string)

	// html模板文件路径。**代表文件夹，*代表文件。*结尾。
	conf["View_Pattern"] = pattern
	// 多静态文件的主文件夹。/结尾。
	conf["View_Static"] = static

	return conf
}
