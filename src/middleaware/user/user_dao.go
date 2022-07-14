/*================================================================
*
*  文件名称：user_dao.go
*  创 建 者: mongia
*  创建日期：2021年12月30日
*
================================================================*/

package middleaware

import (
	"database/sql"
	"errors"

	"github.com/didi/gendry/builder"
	"github.com/didi/gendry/scanner"
)

const (
	userTableName string = "pharmacy.user"
)

func init() {
	scanner.SetTagName("json")
}

// UserDao 用户数据库操作
type UserDao struct {
}

// GetUserInfo 获取用户信息
func (*UserDao) GetUserInfo(db *sql.Tx, where map[string]interface{}, selectFields []string) ([]*UserInfo, int64, error) {
	var count int64 = 0
	if nil == db {
		return nil, count, errors.New("sql.Tx object is nil")
	}
	cond, vals, err := builder.BuildSelect(userTableName, where, selectFields)
	if nil != err {
		return nil, count, err
	}
	row, err := db.Query(cond, vals...)
	if nil != err || nil == row {
		return nil, count, err
	}
	defer row.Close()

	var users []*UserInfo
	err = scanner.Scan(row, &users)
	if nil != err {
		return nil, count, err
	}

	if nil != where["_limit"] {
		res, err := db.Query("SELECT FOUND_ROWS() as total;")
		if nil != err {
			return nil, count, err
		}
		defer res.Close()
		result, err := scanner.ScanMap(res)
		if nil != err {
			return nil, count, err
		}
		if 0 == len(result) {
			return nil, count, errors.New("SELECT FOUND_ROWS() return array 0")
		}
		count = result[0]["count"].(int64)
	}

	return users, count, nil
}

// DeleteUserInfo delete rows from db
func (*UserDao) DeleteUserInfo(db *sql.Tx, where map[string]interface{}) ([]*UserInfo, error) {
	if nil == db {
		return nil, errors.New("sql.Tx object is nil")
	}
	cond, vals, err := builder.BuildSelect(userTableName, where, nil)
	if nil != err {
		return nil, err
	}
	row, err := db.Query(cond, vals...)
	if nil != err || nil == row {
		return nil, err
	}
	defer row.Close()

	var users []*UserInfo
	err = scanner.Scan(row, &users)
	if nil != err {
		return nil, err
	}

	if 0 == len(users) {
		return nil, errors.New("[db] no data")
	}

	cond, vals, err = builder.BuildDelete(userTableName, where)
	if nil != err {
		return nil, err
	}
	_, err = db.Exec(cond, vals...)
	if nil != err {
		return nil, err
	}

	return users, nil
}

// CreateUserInfo 创建用户
func (*UserDao) CreateUserInfo(db *sql.DB, data []map[string]interface{}) (int64, error) {
	if nil == db {
		return 0, errors.New("sql.Tx object is nil")
	}
	cond, vals, err := builder.BuildInsert(userTableName, data)
	if nil != err {
		return 0, err
	}
	result, err := db.Exec(cond, vals...)
	if nil != err || nil == result {
		return 0, err
	}
	return result.LastInsertId()
}

// ModifyUserInfo 修改用户信息
func (*UserDao) ModifyUserInfo(db *sql.Tx, data, where map[string]interface{}) ([]*UserInfo, error) {
	if nil == db {
		return nil, errors.New("sql.Tx object is nil")
	}
	cond, vals, err := builder.BuildSelect(userTableName, where, nil)
	if nil != err {
		return nil, err
	}
	row, err := db.Query(cond, vals...)
	if nil != err || nil == row {
		return nil, err
	}
	defer row.Close()

	var users []*UserInfo
	err = scanner.Scan(row, &users)

	if nil != err {
		return nil, err
	}

	if 0 == len(users) {
		return nil, errors.New("[db] no data")
	}

	cond, vals, err = builder.BuildUpdate(userTableName, where, data)
	if nil != err {
		return nil, err
	}
	result, err := db.Exec(cond, vals...)

	if nil != err {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if nil != err || 0 == rowsAffected {
		return nil, errors.New("[db] no data")
	}

	return users, nil
}
