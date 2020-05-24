package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	Db *gorm.DB
)

func init() {
	host := beego.AppConfig.String("mysql_host")
	port, _ := beego.AppConfig.Int("mysql_port")
	user2 := beego.AppConfig.String("mysql_user")
	password := beego.AppConfig.String("mysql_password")
	db := beego.AppConfig.String("mysql_db")
	prefix := beego.AppConfig.String("mysql_prefix")
	maxIdleConns, _ := beego.AppConfig.Int("mysql_max_idle_conns")
	maxOpenConns, _ := beego.AppConfig.Int("mysql_max_open_conns")

	mysqlConf := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user2, password, host, port, db)
	var err error
	Db, err = gorm.Open("mysql", mysqlConf)
	if err != nil {
		fmt.Println("connect redis error :", err)
		return
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return prefix + defaultTableName
	}

	if beego.AppConfig.String("runmode") == "dev" {
		//开发模式下打印sql
		Db.LogMode(true)
	}

	Db.DB().SetMaxIdleConns(maxIdleConns) //空闲最大连接数
	Db.DB().SetMaxOpenConns(maxOpenConns) //最大连接数
	Db.SingularTable(true)                //全局禁用表名复数
}
