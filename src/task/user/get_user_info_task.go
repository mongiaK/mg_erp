/*================================================================
*
*  文件名称：get_user_info_task.go
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
	"pharmacyerp/util"

	log "pharmacyerp/log"
)

// GetUserInfoTask 处理获取用户信息的入口
type GetUserInfoTask struct {
	Req *userpb.GetUserInfoReq
	Res *userpb.GetUserInfoRes
	Err *util.Errno
}

func (task *GetUserInfoTask) Run(ctx context.Context) {
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

	var selectField = []string{"SQL_CALC_FOUND_ROWS", "*"}
	users, count, err := dao.GetUserInfo(tx, where, selectField)
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

	tx.Commit()
	task.Res.Data.Userinfos = users
	task.Res.Data.Count = count
}

func (task *GetUserInfoTask) setResult() {
	task.Res.Code = task.Err.Code
	task.Res.Msg = task.Err.Msg

	log.Info("GetUserInfo",
		log.Reflect("Request", task.Req),
		log.Reflect("Response", task.Res))
}

func (task *GetUserInfoTask) checkParam() error {
	return nil
}
