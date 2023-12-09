package setting

import (
	"gin-blog/global"
	"github.com/go-ini/ini"
)

var Cfg *ini.File

func Setup(config string) {
	var err error
	configPath := "conf/" + config
	Cfg, err = ini.Load(configPath)
	if err != nil {
		panic("找不到配置文件")
	}
	mapTo("app", &global.Config.AppConfig)
	mapTo("database", &global.Config.DatabaseConfig)
	mapTo("server", &global.Config.ServerConfig)
}

func mapTo(section string, key interface{}) error {
	err := Cfg.Section(section).MapTo(key)
	if err != nil {
		panic("解析配置文件失败")
	}
	return nil
}
