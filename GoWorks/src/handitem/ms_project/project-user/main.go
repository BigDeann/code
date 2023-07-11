package main

import (
	"github.com/gin-gonic/gin"
	srv "test.com/project-common"
	"test.com/project-user/config"
	"test.com/project-user/router"
)

func main() {

	r := gin.Default()
	router.InitRouter(r)
	//grpc服务注册
	grpc := router.RegisterGrpc()
	//grpc服务注册到etcd
	//router.RegisterEtcdServer()
	stop := func() {
		//如果整个服务停了   grpc也就应该停止
		grpc.Stop()
	}
	srv.Run(r, config.C.SC.Name, config.C.SC.Addr, stop)
}
