package mysql

import (
	"database/sql"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

var (
	// TODO 暂时通过全局变量来初始化
	DB *sql.DB
	// 默认 data source name
	dsn = "...(127.0.0.1:3306)/playground?parseTime=true&loc=Local&charset=utf8&collation=utf8_general_ci&interpolateParams=true"
)

func init() {
	// redis 地址可以从环境变量中获取
	mysqlDSN := os.Getenv("POPUPS_MYSQL_DSN")
	if mysqlDSN == "" {
		mysqlDSN = dsn
	}

	var err error
	DB, err = sql.Open("mysql", mysqlDSN)
	if err != nil {
		logrus.WithError(err).Fatal("failed to init mysql")
	}

	DB.SetConnMaxLifetime(5 * time.Minute)
	DB.SetMaxIdleConns(50)
	DB.SetMaxOpenConns(100)

	// Open doesn't open a connection. Validate DSN data:
	err = DB.Ping()
	if err != nil {
		logrus.WithError(err).Fatal("failed to ping mysql")
	}
}
