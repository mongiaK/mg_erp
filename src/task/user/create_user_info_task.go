/*================================================================
*
*  文件名称：create_user_info_task.go
*  创 建 者: mongia
*  创建日期：2022年01月05日
*
================================================================*/

package task

import (
	"context"
	"errors"

	log "pharmacyerp/log"

	middleaware "pharmacyerp/middleaware/user"
	userpb "pharmacyerp/pb/user"
	third "pharmacyerp/third"
	util "pharmacyerp/util"
)

// CreateUserInfoTask 处理创建用户信息的入口
type CreateUserInfoTask struct {
	Req *userpb.CreateUserInfoReq
	Res *userpb.CreateUserInfoRes
	Err *util.Errno
}

// Run task call for outside
func (task *CreateUserInfoTask) Run(ctx context.Context) {
	defer task.setResult()
	if err := task.checkParam(); nil != err {
		log.Error(err.Error())
		task.Err = util.EParam
		return
	}

	var sqldata []map[string]interface{}
	sqldata = append(sqldata, map[string]interface{}{
		"username":  task.Req.Username,
		"password":  task.Req.Password,
		"telephone": task.Req.Telephone,
		"salt":      task.Req.Salt,
		"nickname":  task.Req.Nickname,
		"sex":       task.Req.Sex,
		"born_date": task.Req.BornDate,
		"icon":      task.Req.Icon,
		"card_type": task.Req.CardType,
		"card":      task.Req.Card,
	})

	dao := &middleaware.UserDao{}
	db := third.GetMysqlDB()

	userid, err := dao.CreateUserInfo(db, sqldata)
	if nil != err {
		log.Error(err.Error())
		task.Err = util.EDB
		return
	}

	task.Res.Data = &userpb.UserInfoDB{
		UserId: userid,
	}
}

func (task *CreateUserInfoTask) checkParam() error {
	if "" == task.Req.Username || len(task.Req.Username) > 30 ||
		"" == task.Req.Password || len(task.Req.Password) > 64 ||
		("" != task.Req.Telephone && len(task.Req.Telephone) > 11) ||
		("" != task.Req.Nickname && len(task.Req.Nickname) > 30) ||
		("" != task.Req.Icon && len(task.Req.Icon) > 128) ||
		("" != task.Req.Card && len(task.Req.Card) > 18) {
		return errors.New("[param] rpc request param error")
	}
	return nil
}

func (task *CreateUserInfoTask) setResult() {
	task.Res.Code = task.Err.Code
	task.Res.Msg = task.Err.Msg

	log.Info("CreateUserInfo",
		log.Reflect("Request", task.Req),
		log.Reflect("Response", task.Res))
}
