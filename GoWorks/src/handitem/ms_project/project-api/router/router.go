package router

import (
	"github.com/gin-gonic/gin"
)

// 定义路由接口  里面有一个方法
type Router interface {
	Route(r *gin.Engine)
}

// 定义结构体
type RegisterRouter struct {
}

// New  得到具体实力对象
func New() *RegisterRouter {
	return &RegisterRouter{}
}

// Route 给RegisterRouter创建一个方法
func (*RegisterRouter) Route(router Router, r *gin.Engine) {
	//这里调用的是上面接口的route  而上面的route又可以被别的结构体（实现该接口）对象调用这个Route方法
	router.Route(r)
}

// 因为
var routers []Router

func InitRouter(r *gin.Engine) {

	//方法一：
	//rg := New()
	//以后的模块路由在这进行注册    这里的作用是以后可以不用一直new对象
	//直接用rg对象注册路由，只要在里面放入对应的具体的路由对象
	//rg.Route(&user.RouterUser{}, r)

	//方法二：
	//这里的ro就是
	for _, ro := range routers {
		//这里调用的是上面接口的route  而上面的route又可以被别的结构体（实现该接口）对象调用这个Route方法
		//相当于  这里也是调用接口的route   r.POST("/project/login/getCaptcha", h.getCaptcha)  只不过是弄了一个统一的集合
		ro.Route(r)
	}
}

func Register(ro ...Router) {
	routers = append(routers, ro...)
}
