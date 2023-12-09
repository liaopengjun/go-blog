package gorm

import (
	"fmt"
	"gin-blog/global"
	"gin-blog/pkg/logging"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func InitDb() *gorm.DB {
	var dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		global.Config.DatabaseConfig.User,
		global.Config.DatabaseConfig.Password,
		global.Config.DatabaseConfig.Host,
		global.Config.DatabaseConfig.Name,
	)
	logging.Info(dsn)
	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   global.Config.DatabaseConfig.TablePrefix, //表前缀
			SingularTable: true,                                     //禁用表名复数
		},
	}); err != nil {
		logging.Info(fmt.Errorf("数据库连接失败: %s \n", err))

		panic(fmt.Errorf("数据库连接失败: %s \n", err))
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(10)  // 设置空闲连接池中连接的最大数量
		sqlDB.SetMaxOpenConns(100) // 设置打开数据库连接的最大数量。
		return db
	}

}
