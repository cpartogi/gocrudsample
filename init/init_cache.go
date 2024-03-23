package init

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

func NewRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", viper.GetString(`cache.redis.host`), viper.GetString(`cache.redis.port`)),
		Password: viper.GetString(`cache.redis.password`),
		DB:       0,
	})
	return client
}
