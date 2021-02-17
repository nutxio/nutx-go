package g

import (
	"github.com/go-redis/redis"
	"github.com/nuxtio/nutx-go/config"
	"gorm.io/gorm"
)

var (
	DB    *gorm.DB      // gorm 全局常量
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
	// Oss
	Oss struct {
		AppId      string `mapstructure:"app_id" yaml:"app_id"`
		AppSecret  string `mapstructure:"app_secret" yaml:"app_secret"`
		BucketName string `mapstructure:"bucket_name" yaml:"bucket_name"`
		Endpoint   string `mapstructure:"endpoint" yaml:"endpoint"`
	}
	// 微信配置
	Wx struct {
		// 微信公众号配置
		Mp struct {
			AppID          string `mapstructure:"app_id" yaml:"app_id"`                     //appid
			AppSecret      string `mapstructure:"app_secret" yaml:"app_secret"`             //appsecret
			Token          string `mapstructure:"token" yaml:"token"`                       //token
			EncodingAESKey string `mapstructure:"encoding_aes_key" yaml:"encoding_aes_key"` //EncodingAESKey
			Cache          string `mapstructure:"cache" yaml:"cache"`
		}

		// 微信支付配置
		Pay struct {
			AppID     string `mapstructure:"app_id" yaml:"app_id"`
			MchID     string `mapstructure:"mch_id" yaml:"mch_id"`
			Key       string `mapstructure:"key" yaml:"key"`
			NotifyURL string `mapstructure:"notify_url" yaml:"notify_url"`
		}

		// 微信开放平台配置
		Open struct {
			AppID          string `mapstructure:"app_id" yaml:"app_id"`                     //appid
			AppSecret      string `mapstructure:"app_secret" yaml:"app_secret"`             //appsecret
			Token          string `mapstructure:"token" yaml:"token"`                       //token
			EncodingAESKey string `mapstructure:"encoding_aes_key" yaml:"encoding_aes_key"` //EncodingAESKey
			Cache          string `mapstructure:"cache" yaml:"cache"`
		}

		// 微信小程序配置
		Miniapp struct {
			AppID     string `mapstructure:"app_id" yaml:"app_id"`         //appid
			AppSecret string `mapstructure:"app_secret" yaml:"app_secret"` //appsecret
			Cache     string `mapstructure:"cache" yaml:"cache"`
		}
	}
}

func init() {
	_ = config.VP.Unmarshal(&Config)
}
