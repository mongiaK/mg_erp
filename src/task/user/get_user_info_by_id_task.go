/*================================================================
*
*  文件名称：get_user_info_by_id_task.go
*  创 建 者: mongia
*  创建日期：2022年01月10日
*
================================================================*/

package task

import (
	"context"
	"errors"
	"strconv"

	userpb "pharmacyerp/pb/user"
	"pharmacyerp/util"

	log "pharmacyerp/log"
)

// GetUserInfoByIDTask 处理获取用户信息的入口
type GetUserInfoByIDTask struct {
	Req *userpb.GetUserInfoByIDReq
	Res *userpb.GetUserInfoByIDRes
	Err *util.Errno
}

func (task *GetUserInfoByIDTask) Run(ctx context.Context) {
	defer task.setResult()
	//	key, keyType, err := task.checkParam()
	//	if nil != err {
	//		log.Error(err.Error())
	//		task.Err = util.EParam
	//		return
	//	}
	//
	//	usercache := &middleaware.UserCache{}
	//	user, err := usercache.GetInCache(key, int(keyType))
	//	if nil != err {
	//		log.Warn(err.Error())
	//	} else {
	//		task.Res.Data = user
	//		return
	//	}
	//
	//	where := map[string]interface{}{}
	//	switch keyType {
	//	case userpb.UserItem_USER_ID:
	//		where["user_id"] = key
	//		break
	//	case userpb.UserItem_TELEPHONE:
	//		where["telephone"] = key
	//		break
	//	case userpb.UserItem_CARD:
	//		where["card"] = key
	//		break
	//	case userpb.UserItem_USERNAME:
	//		where["username"] = key
	//		break
	//	}
	//
	//	db := third.GetMysqlDB()
	//	dao := &middleaware.UserDao{}
	//
	//	users, _, err := dao.GetUserInfo(db, where, nil)
	//	if nil != err {
	//		log.Error(err.Error())
	//		task.Err = util.EDB
	//		return
	//	}
	//	task.Res.Data = user
	//
	//	err = usercache.SetInCache(user)
	//	if nil != err {
	//		log.Warn(err.Error())
	//	}
}

func (task *GetUserInfoByIDTask) setResult() {
	task.Res.Code = task.Err.Code
	task.Res.Msg = task.Err.Msg

	log.Info("GetUserInfoByID",
		log.Reflect("Request", task.Req),
		log.Reflect("Response", task.Res))
}

func (task *GetUserInfoByIDTask) checkParam() (string, userpb.UserItem, error) {
	if task.Req.UserId != 0 {
		return strconv.FormatInt(task.Req.UserId, 10), userpb.UserItem_USER_ID, nil
	}
	if len(task.Req.Telephone) != 0 {
		return task.Req.Telephone, userpb.UserItem_TELEPHONE, nil
	}
	if len(task.Req.Username) != 0 {
		return task.Req.Username, userpb.UserItem_USERNAME, nil
	}
	if len(task.Req.Card) != 0 {
		return task.Req.Card, userpb.UserItem_CARD, nil
	}
	return "", userpb.UserItem_ALL, errors.New("[param] param error")
}
