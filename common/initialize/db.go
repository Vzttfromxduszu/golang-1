package initialize

import (
	"fmt"

	"github.com/Vzttfromxduszu/golang-1.git/common/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func MySQL() {
	m := global.Config.Mysql
	var dsn = fmt.Sprintf("%s:%s@%s", m.Username, m.Password, m.Url)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		panic(fmt.Errorf("MySQL启动异常 %w", err))
	}
	sqlDB, err := db.DB() // 获取底层 sql.DB 对象
	if err != nil {
		panic(fmt.Errorf("MySQL启动异常 %w", err))
	}
	sqlDB.SetMaxIdleConns(10)  // SetMaxIdleConns 用于设置连接池中空闲连接的最大数量
	sqlDB.SetMaxOpenConns(100) // SetMaxOpenConns 设置打开数据库连接的最大数量
	global.Db = db

}
