package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"

	"xorm.io/xorm"

	"passport/util"
)

var MasterDB *xorm.Engine

var dataSourceName string

func init() {

	mysqlConfig, _ := util.ConfigFile.GetSection("mysql")

	fillDataSourceName(mysqlConfig)

	if err := initEngine(); err != nil {
		panic(err)
	}

	// 测试数据库连接是否 OK
	if err := MasterDB.Ping(); err != nil {
		panic(err)
	}
}

func fillDataSourceName(mysqlConfig map[string]string) {
	dataSourceName = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		mysqlConfig["user"],
		mysqlConfig["password"],
		mysqlConfig["host"],
		mysqlConfig["port"],
		mysqlConfig["dbname"],
		mysqlConfig["charset"])
}

func initEngine() error {
	var err error

	MasterDB, err = xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		return err
	}

	maxIdle := util.ConfigFile.MustInt("mysql", "max_idle", 2)
	maxConn := util.ConfigFile.MustInt("mysql", "max_conn", 10)

	MasterDB.SetMaxIdleConns(maxIdle)
	MasterDB.SetMaxOpenConns(maxConn)

	MasterDB.ShowSQL(true)

	return nil
}

