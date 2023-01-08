package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var DB *gorm.DB

// Database 连接数据库
func Database(connstring string) {
	db, err := gorm.Open("mysql", connstring)
	if err != nil {
		fmt.Println(err)
		panic("Mysql数据库连接错误")
	}
	fmt.Println("数据库连接成功！！！！！！！")
	//是否打印日志
	db.LogMode(true)
	if gin.Mode() == "release" {
		db.LogMode(false)
	}
	//gorm默认使用复数映射，go代码的单数、复数struct形式都匹配到复数表中：创建表、添加数据时都是如此。
	//指定了db.SingularTable(true)之后，进行严格匹配。
	db.SingularTable(true)
	//配置数据库连接池
	//设置池中空闲连接数的上限。
	db.DB().SetMaxIdleConns(5)
	//设置池中“打开”连接(使用中+空闲连接)数量的上限。
	db.DB().SetMaxOpenConns(10)
	// 设置最大生存时间为30秒
	db.DB().SetConnMaxLifetime(time.Second * 30)
	DB = db
	//数据迁移
	migration()
}
