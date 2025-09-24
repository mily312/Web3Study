package conf

import (
	"fmt"

	"github.com/spf13/viper"
)

// 读取配置文件信息并初始化
func InitConfig() {
	viper.SetConfigName("settings")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./conf/")
	err := viper.ReadInConfig()
	if err != nil {
		panic("viper read conf file error:" + err.Error())
	}

	fmt.Println("端口号：" + viper.GetString("server.port"))
}
