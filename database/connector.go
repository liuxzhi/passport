package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var MasterDB *xorm.Engine

var dns string

func init() {
	dns = fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=%s&parseTime=True&loc=Local",
		"root",
		"123456",
		"172.16.113.119",
		"3306",
		"utf8")
	if err := initEngine(); err != nil {
		panic(err)
	}

	// 测试数据库连接是否 OK
	if err := MasterDB.Ping(); err != nil {
		panic(err)
	}
}


func initEngine() error {
	var err error

	MasterDB, err = xorm.NewEngine("mysql", dns)
	if err != nil {
		return err
	}

	maxIdle := 2
	maxConn := 10

	MasterDB.SetMaxIdleConns(maxIdle)
	MasterDB.SetMaxOpenConns(maxConn)

	MasterDB.ShowSQL(true)

	return nil
}

