/*================================================================
*
*  文件名称：delete_user_info_task.go
*  创 建 者: mongia
*  创建日期：2022年01月05日
*
================================================================*/

package task

import (
	"context"
	middleaware "pharmacyerp/middleaware/user"
	userpb "pharmacyerp/pb/user"
	third "pharmacyerp/third"
	util "pharmacyerp/util"

	log "pharmacyerp/log"
)

// DeleteUserInfoTask 处理删除用户信息的入口
type DeleteUserInfoTask struct {
	Req *userpb.DeleteUserInfoReq
	Res *userpb.DeleteUserInfoRes
	Err *util.Errno
}

func (task *DeleteUserInfoTask) Run(ctx context.Context) {
	defer task.setResult()
	if err := task.checkParam(); nil != err {
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
	if nil != err {
		log.Error(err.Error())
		task.Err = util.EDB
		return
	}
	users, err := dao.DeleteUserInfo(tx, where)
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

func (task *DeleteUserInfoTask) setResult() {
	task.Res.Code = task.Err.Code
	task.Res.Msg = task.Err.Msg

	log.Info("DeleteUserInfo",
		log.Reflect("Request", task.Req),
		log.Reflect("Response", task.Res))
}

func (task *DeleteUserInfoTask) checkParam() error {
	return nil
}
