package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	common "test.com/project-common"
	"test.com/project-common/errs"
	loginServiceV1 "test.com/project-user/pkg/service/login.service.v1"
	"time"
)

// 有具体的缓存实例对象字段    因为这里要有缓存的操作  所以有缓存的字段
type HandlerUser struct {
}

// 得到具体的api实例对象
func New() *HandlerUser {
	return &HandlerUser{}
}

// 获取验证码的接口
func (*HandlerUser) getCaptcha(c *gin.Context) {
	result := &common.Result{}
	mobile := c.PostForm("mobile")
	//拿到对应上面文  为后面调用传递参数
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	//执行rpc调用（这个方法在服务端实现并返回结果）
	rsp, err := LoginServiceClient.GetCaptcha(ctx, &loginServiceV1.CaptchaMessage{
		Mobile: mobile,
	})
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}
	c.JSON(http.StatusOK, result.Success(rsp.Code))
}
