/*================================================================
*
*  文件名称：modify_user_info_task.go
*  创 建 者: mongia
*  创建日期：2022年01月05日
*
================================================================*/

package task

import (
	"context"
	"errors"
	"strings"

	log "pharmacyerp/log"
	middleaware "pharmacyerp/middleaware/user"
	userpb "pharmacyerp/pb/user"
	third "pharmacyerp/third"
	util "pharmacyerp/util"
)

// ModifyUserInfoTask 处理修改用户信息的入口
type ModifyUserInfoTask struct {
	Req *userpb.ModifyUserInfoReq
	Res *userpb.ModifyUserInfoRes
	Err *util.Errno
}

func (task *ModifyUserInfoTask) Run(ctx context.Context) {
	defer task.setResult()
	sqldata, err := task.checkParam()
	if nil != err {
		log.Error(err.Error())
		task.Err = util.EParam
		return
	}

	where, err := generateSQLWhere(task.Req.Cond)
	if nil != err {
		log.Error(err.Error())
		task.Err = util.EParam
		return
	}

	dao := &middleaware.UserDao{}
	db := third.GetMysqlDB()

	tx, err := db.Begin()

	users, err := dao.ModifyUserInfo(tx, sqldata, where)
	if nil != err {
		log.Error(err.Error())
		task.Err = util.EDB
		tx.Rollback()
		return
	}

	if nil == users {
		task.Err = util.EDBNoData
		tx.Rollback()
		return
	}

	usercache := &middleaware.UserCache{}
	err = usercache.RemoveInCache(users)
	if nil != err {
		log.Error(err.Error())
		task.Err = util.ERedis
		tx.Rollback()
		return
	}
	tx.Commit()
}

func (task *ModifyUserInfoTask) checkParam() (map[string]interface{}, error) {
	data := make(map[string]interface{}, 0)
	for key, val := range userpb.UserItem_name {
		val = strings.ToLower(val)
		switch key & int32(task.Req.ModifyItems) {
		case int32(userpb.UserItem_USERNAME):
			if "" == task.Req.Username || len(task.Req.Username) > 30 {
				return nil, errors.New("[param] username irregular")
			}
			data[val] = task.Req.Username
			break
		case int32(userpb.UserItem_PASSWORD):
			if "" == task.Req.Password || len(task.Req.Password) > 64 {
				return nil, errors.New("[param] password irregular")
			}
			data[val] = task.Req.Password
			break
		case int32(userpb.UserItem_TELEPHONE):
			if "" != task.Req.Telephone && len(task.Req.Telephone) > 11 {
				return nil, errors.New("[param] telephone irregular")
			}
			data[val] = task.Req.Telephone
			break
		case int32(userpb.UserItem_NICKNAME):
			if "" != task.Req.Nickname && len(task.Req.Nickname) > 30 {
				return nil, errors.New("[param] nickname irregular")
			}
			data[val] = task.Req.Nickname
			break
		case int32(userpb.UserItem_SEX):
			if task.Req.Sex > 1 {
				return nil, errors.New("[param] sex irregular")
			}
			data[val] = task.Req.Sex
			break
		case int32(userpb.UserItem_SALT):
			if "" != task.Req.Salt && len(task.Req.Salt) > 6 {
				return nil, errors.New("[param] salt irregular")
			}
			data[val] = task.Req.Salt
			break
		case int32(userpb.UserItem_BORN_DATE):
			if 0 == task.Req.BornDate {
				return nil, errors.New("[param] born_date irregular")
			}
			data[val] = task.Req.BornDate
			break
		case int32(userpb.UserItem_ICON):
			if "" != task.Req.Icon && len(task.Req.Icon) > 128 {
				return nil, errors.New("[param] icon irregular")
			}
			data[val] = task.Req.Icon
			break
		case int32(userpb.UserItem_CARD_TYPE):
			data[val] = task.Req.CardType
			break
		case int32(userpb.UserItem_STATUS):
			data[val] = task.Req.Status
			break
		case int32(userpb.UserItem_CARD):
			if "" != task.Req.Card && len(task.Req.Card) > 18 {
				return nil, errors.New("[param] card irregular")
			}
			data[val] = task.Req.Card
			break
		}
	}

	if len(data) == 0 {
		return nil, errors.New("[param] no item modify")
	}

	return data, nil
}

func (task *ModifyUserInfoTask) setResult() {
	task.Res.Code = task.Err.Code
	task.Res.Msg = task.Err.Msg

	log.Info("ModifyUserInfo",
		log.Reflect("Request", task.Req),
		log.Reflect("Response", task.Res))
}
