package g

import (
	"github.com/go-redis/redis"
	"github.com:nuxtio/nutx-go/config"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB         // gorm 全局常量
	Redis *redis.Client // redis 全局常量
)

// 配置相关
var Config struct {
	// 端口
	Port string `mapstructure:"port" yaml:"port"`
	// 数据库相关配置
	Database struct {
		Dsn string `mapstructure:"dsn" yaml:"dsn"`
	}
	// Redis
	Redis struct {
		Host        string `mapstructure:"host" yaml:"host"`
		Password    string `mapstructure:"password" yaml:"password"`
		Database    int    `mapstructure:"database" yaml:"database"`
		MaxIdle     int    `mapstructure:"max_idle" yaml:"max_idle"`
		MaxActive   int    `mapstructure:"max_active" yaml:"max_active"`
		IdleTimeout int    `mapstructure:"idle_timeout" yaml:"idle_timeout"`
	}
}

func init() {
	_ = config.VP.Unmarshal(&Config)
}
