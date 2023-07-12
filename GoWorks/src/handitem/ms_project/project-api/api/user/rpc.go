package user

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	loginServiceV1 "test.com/project-user/pkg/service/login.service.v1"
)

var LoginServiceClient loginServiceV1.LoginServiceClient

// 拿到rpc的对象
func InitRpcUserClient() {

	//引入etcd
	//etcdRegister := discovery.NewResolver(config.C.EtcdConfig.Addrs, logs.LG)
	//resolver.Register(etcdRegister)
	//conn, err := grpc.Dial("etcd:///user", grpc.WithTransportCredentials(insecure.NewCredentials()))

	//1、连接到server端口,此处禁用安全连接，没有加密和验证  拿到连接对象
	conn, err := grpc.Dial("127.0.0.1:8881", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	//2、建立连接
	LoginServiceClient = loginServiceV1.NewLoginServiceClient(conn)

	//3、调用方法：这里的LoginServiceClient就可以直接访问到另一端的方法了
	//LoginServiceClient.GetCaptcha()
}
