/*================================================================
*
*  文件名称：user_dao.go
*  创 建 者: mongia
*  创建日期：2021年12月30日
*
================================================================*/

package middleaware

import (
	"context"
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

type UserDao struct {
}

// GetUserInfo 获取用户信息
func (*UserDao) GetUserInfo(db *sql.Tx, where map[string]interface{}, selectFields []string) ([]*UserInfo, error) {
	users, err := getUserMulti(db, where, selectFields)
	if nil != err {
		return nil, err
	}

	return users, nil
}

// DeleteUserInfo delete rows from db
func (*UserDao) DeleteUserInfo(db *sql.Tx, where map[string]interface{}) ([]*UserInfo, error) {
	users, err := getUserMulti(db, where, nil)
	if nil != err {
		return nil, err
	}

	if 0 == len(users) {
		return nil, errors.New("[db] no data")
	}

	_, err = DeleteUser(db, where, nil)
	if nil != err {
		return nil, err
	}

	return users, nil
}

// CreateUserInfo 创建用户
func (*UserDao) CreateUserInfo(db *sql.DB, data []map[string]interface{}) (int64, error) {
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
	users, err := getUserMulti(db, where, nil)
	if nil != err {
		return nil, err
	}

	if 0 == len(users) {
		return nil, errors.New("[db] no data")
	}

	rowsAffected, err := UpdateUser(db, where, data)
	if nil != err {
		return nil, err
	}

	if 0 == rowsAffected {
		return nil, errors.New("[db] no data")
	}

	return users, nil
}

func (*UserDao) GetUserCount(db *sql.DB, where map[string]interface{}, selectFields []string) (int64, error) {
	if nil == db {
		return 0, errors.New("sql.Tx object couldn't be nil")
	}
	result, err := builder.AggregateQuery(context.Background(), db, userTableName, where, builder.AggregateCount("*"))
	if nil != err {
		return 0, err
	}

	return result.Int64(), err
}

//GetOne gets one record from table user by condition "where"
func (*UserDao) GetUserOne(db *sql.DB, where map[string]interface{}, selectFields []string) (*UserInfo, error) {
	if nil == db {
		return nil, errors.New("sql.DB object couldn't be nil")
	}
	cond, vals, err := builder.BuildSelect(userTableName, where, selectFields)
	if nil != err {
		return nil, err
	}
	row, err := db.Query(cond, vals...)
	if nil != err || nil == row {
		return nil, err
	}
	defer row.Close()
	var res *UserInfo
	err = scanner.Scan(row, &res)
	return res, err
}

//GetMulti gets multiple records from table user by condition "where"
func getUserMulti(db *sql.Tx, where map[string]interface{}, selectFields []string) ([]*UserInfo, error) {
	if nil == db {
		return nil, errors.New("sql.Tx object couldn't be nil")
	}
	cond, vals, err := builder.BuildSelect(userTableName, where, selectFields)
	if nil != err {
		return nil, err
	}
	row, err := db.Query(cond, vals...)
	if nil != err || nil == row {
		return nil, err
	}
	defer row.Close()

	var res []*UserInfo
	err = scanner.Scan(row, &res)
	return res, err
}

//UpdateUser updates the table user
func UpdateUser(db *sql.Tx, where, data map[string]interface{}) (int64, error) {
	if nil == db {
		return 0, errors.New("sql.Tx object couldn't be nil")
	}
	cond, vals, err := builder.BuildUpdate(userTableName, where, data)
	if nil != err {
		return 0, err
	}
	result, err := db.Exec(cond, vals...)
	if nil != err {
		return 0, err
	}
	return result.RowsAffected()
}

// DeleteUser deletes matched records in user
func DeleteUser(db *sql.Tx, where, data map[string]interface{}) (int64, error) {
	if nil == db {
		return 0, errors.New("sql.Tx object couldn't be nil")
	}
	cond, vals, err := builder.BuildDelete(userTableName, where)
	if nil != err {
		return 0, err
	}
	result, err := db.Exec(cond, vals...)
	if nil != err {
		return 0, err
	}
	return result.RowsAffected()
}
