package wechat

import (
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram"
	ma "github.com/silenceper/wechat/v2/miniprogram/config"
	"github.com/silenceper/wechat/v2/officialaccount"
	mp "github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/openplatform"
	open "github.com/silenceper/wechat/v2/openplatform/config"
	"github.com/silenceper/wechat/v2/pay"
	payConfig "github.com/silenceper/wechat/v2/pay/config"
	"github.com:nuxtio/nutx-go/g"
)

// 获取公众号服务
func GetOfficialAccount() *officialaccount.OfficialAccount {
	wx := wechat.NewWechat()
	c := getCache(g.Config.Wx.Mp.Cache)
	cfg := &mp.Config{
		AppID:          g.Config.Wx.Mp.AppID,
		AppSecret:      g.Config.Wx.Mp.AppSecret,
		Token:          g.Config.Wx.Mp.Token,
		EncodingAESKey: g.Config.Wx.Mp.EncodingAESKey,
		Cache:          c,
	}
	return wx.GetOfficialAccount(cfg)
}

// 获取支付服务
func GetPay() *pay.Pay {
	wx := wechat.NewWechat()
	cfg := &payConfig.Config{
		AppID:     g.Config.Wx.Pay.AppID,
		MchID:     g.Config.Wx.Pay.MchID,
		Key:       g.Config.Wx.Pay.Key,
		NotifyURL: g.Config.Wx.Pay.NotifyURL,
	}
	return wx.GetPay(cfg)
}

// 获取小程序服务
func GetMiniProgram() *miniprogram.MiniProgram {
	wx := wechat.NewWechat()
	c := getCache(g.Config.Wx.Mp.Cache)
	cfg := &ma.Config{
		AppID:     g.Config.Wx.Miniapp.AppID,
		AppSecret: g.Config.Wx.Miniapp.AppSecret,
		Cache:     c,
	}
	return wx.GetMiniProgram(cfg)
}

// 获取开放平台
func GetOpenPlatform() *openplatform.OpenPlatform {
	wx := wechat.NewWechat()
	c := getCache(g.Config.Wx.Mp.Cache)
	cfg := &open.Config{
		AppID:          g.Config.Wx.Open.AppID,
		AppSecret:      g.Config.Wx.Open.AppSecret,
		Token:          g.Config.Wx.Open.Token,
		EncodingAESKey: g.Config.Wx.Open.EncodingAESKey,
		Cache:          c,
	}
	return wx.GetOpenPlatform(cfg)
}

func getCache(config string) cache.Cache {
	switch config {
	case "redis":
		opts := &cache.RedisOpts{
			Host:        g.Config.Redis.Host,
			Password:    g.Config.Redis.Password,
			Database:    g.Config.Redis.Database,
			MaxIdle:     g.Config.Redis.MaxIdle,
			MaxActive:   g.Config.Redis.MaxActive,
			IdleTimeout: g.Config.Redis.IdleTimeout,
		}
		return cache.NewRedis(opts)
	default:
		return cache.NewMemory()
	}
}
