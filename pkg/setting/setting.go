package setting

import (
	"log"

	"github.com/spf13/viper"
)

var (
	//AppConfig 配置
	AppConfig *viper.Viper
)

func init() {
	AppConfig = viper.New()

	//设置配置文件的名字
	AppConfig.SetConfigName("config")

	//添加配置文件所在的路径,注意在Linux环境下%GOPATH要替换为$GOPATH
	AppConfig.AddConfigPath("conf/")
	AppConfig.AddConfigPath("./")

	//设置配置文件类型
	AppConfig.SetConfigType("yaml")

	if err := AppConfig.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatalln("未找到配置文件")
		} else {
			log.Fatalln("配置文件读取失败:" + err.Error())
		}
	}

	//	log.Printf("age: %s, name: %s \n", AppConfig.Get("information.age"), AppConfig.Get("information.name"))
}
