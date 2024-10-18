package config

import (
	tool "github.com/lie-flat-planet/service-init-tool"
	"github.com/lie-flat-planet/service-init-tool/component/mysql"
	"github.com/lie-flat-planet/service-init-tool/component/redis"
)

func init() {
	if err := tool.Init("./", &Config); err != nil {
		panic(err)
	}
}

var Config = struct {
	Server *tool.Server
	Mysql  *mysql.Mysql
	Redis  *redis.Redis
}{
	Server: &tool.Server{
		Name:     "gin-demo",
		Code:     333 * 1e3,
		HttpPort: 8081,
		RunMode:  "debug",
	},
	Mysql: &mysql.Mysql{
		Config: mysql.Config{
			Host:        "127.0.0.1:3306",
			User:        "root",
			Password:    "",
			DbName:      "",
			MaxIdleConn: 1,
			MaxOpenConn: 2,
		},
	},
	Redis: &redis.Redis{
		Config: redis.Config{
			Host:     "127.0.0.1:6379",
			Password: "",
			DB:       0,
			PoolSize: 5,
			Timeout:  0,
		},
	},
}