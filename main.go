package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"go_gin_example/models"
	"go_gin_example/routers"
	"go_gin_example/utils/conf"
	"go_gin_example/utils/redis"
	"log"
	"net/http"
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

	gin.SetMode(conf.ServiceConf.RunMode)


	router := routers.Router()
	readTimeout := conf.ServiceConf.ReadTimeout
	writeTimeout := conf.ServiceConf.WriteTimeout
	endPoint := fmt.Sprintf("%s:%d", conf.ServiceConf.HttpHost,conf.ServiceConf.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        router,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	err := server.ListenAndServe()
	if err != nil{
		log.Fatalf("err:%v", err)
	}

}
