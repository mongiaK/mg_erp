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

	if 0 != task.Req.Cond.Pagesize {
		count, err := dao.GetUserCount(db, where, nil)
		if nil != err {
			log.Error(err.Error())
			task.Err = util.EDB
			return
		}

		task.Res.Data.Count = count
		if count == 0 {
			return
		}

		where["_limit"] = []uint{uint(task.Req.Cond.Pagenum * task.Req.Cond.Pagesize), uint(task.Req.Cond.Pagesize)}
	}

	tx, err := db.Begin()

	users, err := dao.GetUserInfo(tx, where, nil)
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
