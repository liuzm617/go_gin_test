package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"go_gin_example/models"
	"go_gin_example/routers"
	"go_gin_example/utils/conf"
	"go_gin_example/utils/logging"
	"go_gin_example/utils/redis"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

//func init(){
//	conf.Init()
//}


// @title Golang Gin API
// @version 1.0
// @description An example of gin
// @termsOfService
// @contact.name API Support
// @contact.email
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host
// @BasePath /
func main() {
	flag.Parse()
	// 初始配置
	conf.Init()

	// 初始化mysql连接
	models.Init()
	defer models.Close()

	//初始化redis
	redis.Init()

	// 初始化log
	logging.Init()
	gin.SetMode(conf.ServiceConf.RunMode)


	router := routers.Router()
	readTimeout := conf.ServiceConf.ReadTimeout
	writeTimeout := conf.ServiceConf.WriteTimeout
	endPoint := fmt.Sprintf("%s:%d", conf.ServiceConf.HttpHost,conf.ServiceConf.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        router,
		ReadTimeout:    readTimeout * time.Second,
		WriteTimeout:   writeTimeout * time.Second,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	go func(){
		err := server.ListenAndServe()
		if err != nil{
			log.Fatalf("err:%v", err)
		}
	}()
	//
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<- quit
	log.Println("Shutdown Server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil{
		log.Fatal("shutdown server err", err)
	}
	log.Println("Server exiting")
}
