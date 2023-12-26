package models

import (
	"fmt"
	"gopkg.in/ini.v1"
	"gorm.io/gorm/logger"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

var err error

func init() {
	// 读取conf/app.ini 中的mysql配置
	config, err := ini.Load("./conf/app.ini")
	if err != nil {
		log.Fatal("数据库配置文件错误")
	}
	ip := config.Section("mysql").Key("ip").String()
	port := config.Section("mysql").Key("port").String()
	user := config.Section("mysql").Key("user").String()
	password := config.Section("mysql").Key("password").String()
	database := config.Section("mysql").Key("database").String()
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, ip, port, database)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		QueryFields: true,
		Logger:      logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("数据库连接错误, %s", err)
	}
}
