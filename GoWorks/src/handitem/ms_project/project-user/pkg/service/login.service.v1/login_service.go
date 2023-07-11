package login_service_v1

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"log"
	common "test.com/project-common"
	"test.com/project-common/errs"
	"test.com/project-user/pkg/dao"
	"test.com/project-user/pkg/model"
	"test.com/project-user/pkg/repo"
	"time"
)

// grpc服务的结构体
type LoginService struct {
	UnimplementedLoginServiceServer
	Cache repo.Cache
}

func New() *LoginService {
	return &LoginService{
		Cache: dao.Rc,
	}
}

// 实现grpc具体的业务逻辑
func (ls *LoginService) GetCaptcha(c context.Context, msg *CaptchaMessage) (*CaptchaResponse, error) {
	//获取参数   api层调用的时候会直接传入mobile参数  这里直接拿到就可以了
	mobile := msg.Mobile
	//校验参数
	if !common.VerifyMobile(mobile) {
		return nil, errs.GrpcError(model.NoLegalMobile)
	}

	code := "123456"
	go func() {
		time.Sleep(time.Second * 2)
		zap.L().Info("短信平台调用成功，发送短信 INFO")
		//拿到context 对象并设置两分钟的时间
		c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		//存储验证码到redis中设置过期时间15分钟
		err := ls.Cache.Put(c, mobile, code, time.Minute*15)
		if err != nil {
			fmt.Printf("验证码存入redis出错，cause by%v", err)
		}
		log.Printf("手机号存入redis成功 %s : %s", mobile, code)
	}()
	//
	return &CaptchaResponse{Code: code}, nil
}
