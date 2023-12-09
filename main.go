package main

import (
	"context"
	"flag"
	"fmt"
	"gin-blog/global"
	"gin-blog/pkg/gorm"
	"gin-blog/pkg/setting"
	"gin-blog/routers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var config string

func main() {

	flag.StringVar(&config, "c", "config.ini", "choose config file.")
	flag.Parse()

	// 加载配置
	setting.Setup(config)
	// 连接数据库
	global.Db = gorm.InitDb()

	// 注册路由
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", global.Config.ServerConfig.HttpPort),
		Handler:        router,
		ReadTimeout:    global.Config.ServerConfig.ReadTimeout * time.Millisecond,
		WriteTimeout:   global.Config.ServerConfig.WriteTimeout * time.Millisecond,
		MaxHeaderBytes: 1 << 20,
	}

	// 开启goroutine启动服务
	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Printf("Listen: %s\n", err)
		}
	}()

	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit // 阻塞等待信号

	log.Println("Shutdown Server ...")

	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
