package router

import (
	"BlogSystem/global"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"BlogSystem/middleware"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// 自定义函数类型
type IFnRegistRouter = func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup)

// 遍历路由的切面，对各个模块进行初始化，各个模块的路由都放在这个切面里
var gFnRouters []IFnRegistRouter

func InitRouter() {
	r := gin.Default()

	// 挂载中间件
	r.Use(middleware.Cors())

	//uri前缀
	rgPublic := r.Group("/api/v1/public")
	rgAuth := r.Group("/api/v1")

	rgAuth.Use(middleware.Auth())

	InitBasePlatmRouter()

	for _, fnRegistRouter := range gFnRouters {
		fnRegistRouter(rgPublic, rgAuth)
	}

	var stPort = viper.GetString("server.port")
	if stPort == "" {
		stPort = "8888"
	}

	/* 优雅关闭方式
	 */
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", stPort),
		Handler: r,
	}

	// 启动方式改为srv.ListenAndServe()这种方式
	// err := r.Run(fmt.Sprintf(":%s", stPort))
	// if err != nil {
	// 	panic("server run error:" + err.Error())
	// }

	go func() {
		global.Logger.Info("start server on port:" + stPort)
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.Logger.Error("listen: %s\n", err.Error())
			//log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no params) by default sends syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		global.Logger.Error("Server Shutdown:%s\n", err.Error())
		//log.Println("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}

// 往切面中添加函数类型
func RegistRouter(fn IFnRegistRouter) {
	if fn == nil {
		return
	}

	gFnRouters = append(gFnRouters, fn)
}

// 初始化各个模块的路由
func InitBasePlatmRouter() {
	InitUserRouter()
}
