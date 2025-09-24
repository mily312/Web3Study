package cmd

import (
	"BlogSystem/conf"
	"BlogSystem/global"
	"BlogSystem/router"
	"BlogSystem/utils"
	"fmt"
)

// 项目启动
func Start() {
	//初始化全局错误
	var initError error

	//初始化系统配置文件
	conf.InitConfig()

	//初始化日志组件
	global.Logger = conf.InitLogger()

	//初始化数据库连接
	db, errDB := conf.InitDBConfig()
	global.DB = db
	initError = utils.AppendError(initError, errDB)

	//判断各个模块初始化有没有错误，需要写在初始化路由前面（在初始化路由时，程序会一直执行，不会执行下面的方法）
	if initError != nil {
		if global.Logger != nil {
			global.Logger.Error(initError.Error())
		}

		panic(initError.Error())
	}

	//初始化路由
	router.InitRouter()
}

// 项目结束
func Clean() {
	fmt.Println("=============== server stopped ===============")
}
