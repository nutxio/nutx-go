package initialize

import (
	"fmt"
	"github.com/nuxtio/nutx-go/g"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func InitDB() {
	db, err := gorm.Open(mysql.Open(g.Config.Database.Dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 禁用表复数
		},
	})

	if err != nil {
		panic(fmt.Sprintf("数据库初始化失败：%s", err))
	}

	g.DB = db
}
