package router

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"log"
	"net"
	"test.com/project-common/discovery"
	"test.com/project-common/logs"
	"test.com/project-user/config"
	loginServiceV1 "test.com/project-user/pkg/service/login.service.v1"
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
	//这里调用的是上面的route  而上面的route又可以被别的结构体（实现该接口）对象调用这个Route方法
	router.Route(r)
}

var routers []Router

func InitRouter(r *gin.Engine) {
	//rg := New()
	//以后的模块路由在这进行注册    这里的作用是以后可以不用一直new对象
	//直接用rg对象注册路由，只要在里面放入对应的具体的路由对象
	//rg.Route(&user.RouterUser{}, r)
	for _, ro := range routers {
		//这里调用的是上面的route  而上面的route又可以被别的结构体（实现该接口）对象调用这个Route方法
		//相当于   r.POST("/project/login/getCaptcha", h.getCaptcha)  只不过是弄了一个统一的集合
		ro.Route(r)
	}
}

func Register(ro ...Router) {
	routers = append(routers, ro...)
}

// gRPC的相关配置
type gRPCConfig struct {
	Addr string
	//注册服务的字段
	RegisterFunc func(*grpc.Server)
}

// 注册gRPC  就需要拿到config文件里的对应参数
func RegisterGrpc() *grpc.Server {
	//完成相关配置
	c := gRPCConfig{
		Addr: config.C.GC.Addr,
		RegisterFunc: func(g *grpc.Server) {
			//在grpc服务端去注册我们自己编写的服务 第二个参数是引用  所以客户端可以拿着grpc对象直接调用它的方法
			loginServiceV1.RegisterLoginServiceServer(g, loginServiceV1.New())
		}}
	//  1、创建grpc服务
	s := grpc.NewServer()
	//2、在grpc服务端去注册我们自己编写的服务
	c.RegisterFunc(s)
	//3、开启端口
	lis, err := net.Listen("tcp", c.Addr)
	if err != nil {
		log.Println("cannot listen")
	}
	//如果这里不启动协程  grpc有问题的话  整个项目就不能启动了
	go func() {
		//4、启动服务  将上面的端口放到这来
		err = s.Serve(lis)
		if err != nil {
			log.Println("server started error", err)
			return
		}
	}()
	return s
}

func RegisterEtcdServer() {
	etcdRegister := discovery.NewResolver(config.C.EtcdConfig.Addrs, logs.LG)
	resolver.Register(etcdRegister)
	info := discovery.Server{
		Name:    config.C.GC.Name,
		Addr:    config.C.GC.Addr,
		Version: config.C.GC.Version,
		Weight:  config.C.GC.Weight, //权重  再负载均衡的时候使用的
	}
	r := discovery.NewRegister(config.C.EtcdConfig.Addrs, logs.LG)
	_, err := r.Register(info, 2)
	if err != nil {
		log.Fatalln(err)
	}
}
