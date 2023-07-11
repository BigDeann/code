package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"test.com/project-common/logs"
)

// 在点C的时候触发初始化
var C = InitConfig()

// 一级 总配置结构体
type Config struct {
	viper      *viper.Viper
	SC         *ServerConfig
	EtcdConfig *EtcdConfig
}

// 二级结构体 服务端口的相关配置
type ServerConfig struct {
	Name string
	Addr string
}

// etcd
type EtcdConfig struct {
	Addrs []string
}

func InitConfig() *Config {
	v := viper.New()
	conf := &Config{viper: v}
	//获取当前工作目录的路劲
	workDir, _ := os.Getwd()
	//fmt.Println(workDir) // /Users/dean/GoWorks/src/handitem/ms_project/project-user
	//对应的配置文件的名字
	conf.viper.SetConfigName("app")
	//对应的配置文件类型
	conf.viper.SetConfigType("yaml")
	//在当前工作目录下的config包找
	conf.viper.AddConfigPath(workDir + "/config")
	//尝试进行配置读取  就是说开启读取配置文件功能
	err := conf.viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	//将相关的读操作启动   读完之后  会自动将数据填充到对应的结构体上去
	conf.ReadServerConfig()
	conf.InitZapLog()
	conf.ReadEtcdConfig()
	return conf
}

// 日志打印的相关配置
func (c *Config) InitZapLog() {
	//从配置中读取日志配置，初始化日志   并直接映射到LogConfig的结构体上
	lc := &logs.LogConfig{
		DebugFileName: c.viper.GetString("zap.debugFileName"),
		InfoFileName:  c.viper.GetString("zap.infoFileName"),
		WarnFileName:  c.viper.GetString("zap.warnFileName"),
		MaxSize:       c.viper.GetInt("maxSize"),
		MaxAge:        c.viper.GetInt("maxAge"),
		MaxBackups:    c.viper.GetInt("maxBackups"),
	}
	//把上面读到结构体上的数据放到InitLogger	函数里面进行初始化
	err := logs.InitLogger(lc)
	if err != nil {
		log.Fatalln(err)
	}
}

// 因为这里始终传递的是指针，所以对于我们取的时候才能够 直接取道对应的值
func (c *Config) ReadServerConfig() {
	sc := &ServerConfig{}
	sc.Name = c.viper.GetString("server.name")
	sc.Addr = c.viper.GetString("server.addr")
	c.SC = sc
}

func (c *Config) ReadEtcdConfig() {
	ec := &EtcdConfig{}
	var addrs []string
	//因为可能有多个etcd的端口所以要拿切片接收数据
	err := c.viper.UnmarshalKey("etcd.addrs", &addrs)
	if err != nil {
		log.Fatalln(err)
	}
	ec.Addrs = addrs
	c.EtcdConfig = ec
}
