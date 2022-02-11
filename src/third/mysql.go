/*================================================================
*
*  文件名称：mysql.go
*  创 建 者: mongia
*  创建日期：2021年12月29日
*
================================================================*/

package db

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/didi/gendry/manager"
)

var mysqldb *sql.DB

// MysqlInit is a function for init db
func MysqlInit(database, username, password, host string, port, conns int) error {
	db, err := manager.New(database, username, password, host).Set(
		manager.SetCharset("utf8"),
		manager.SetAllowCleartextPasswords(true),
		manager.SetInterpolateParams(true),
		manager.SetTimeout(1*time.Second),
		manager.SetReadTimeout(1*time.Second),
	).Port(port).Open(true)

	if nil != err {
		return err
	}
	db.SetMaxOpenConns(conns)
	db.SetMaxIdleConns(conns / 2)
	db.SetConnMaxIdleTime(10 * time.Minute)

	mysqldb = db
	return nil
}

// GetMysqlDB 返回数据库对象
func GetMysqlDB() *sql.DB {
	return mysqldb
}

// MysqlUninit call this when exit
func MysqlUninit() {
	mysqldb.Close()
}
