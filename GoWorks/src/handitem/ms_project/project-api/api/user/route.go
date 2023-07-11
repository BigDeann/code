package user

import (
	"github.com/gin-gonic/gin"
	"log"
	"test.com/project-api/router"
)

func init() {
	log.Println("init router user!！")
	//将所有的路由对象放到一个集合里面   对应router包里面的register
	router.Register(&RouterUser{})
}

// 结构体实现了Router接口的方法
type RouterUser struct {
}

// 这里相当于gva的initRouter
func (*RouterUser) Route(r *gin.Engine) {
	InitRpcUserClient()
	//得到调用具体方法的api对象
	h := New()
	//注册路由  并调用具体方法
	r.POST("/project/login/getCaptcha", h.getCaptcha)
}
