package dao

import (
	"context"
	"github.com/go-redis/redis/v8"
	"test.com/project-user/config"
	"time"
)

// 导出经过实例化的redis对象  以后可以通过包名导出来
var Rc *RedisCache

type RedisCache struct {
	rdb *redis.Client
}

// 初始化连接redis  返回redis实例对象
func init() {
	//以后可以在这里更换其他的缓存实例 但是service层里面的代码不需要改动   实现高内聚  低耦合  MongoDB啥的
	rdb := redis.NewClient(config.C.ReadRedisConfig())
	Rc = &RedisCache{
		rdb: rdb,
	}
}

// 实现cache接口的两个方法
func (rc *RedisCache) Put(ctx context.Context, key, value string, expire time.Duration) error {
	err := rc.rdb.Set(ctx, key, value, expire).Err()
	return err
}
func (rc *RedisCache) Get(ctx context.Context, key string) (string, error) {
	result, err := rc.rdb.Get(ctx, key).Result()
	return result, err
}
