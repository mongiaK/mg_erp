/*================================================================
*
*  文件名称：user_service.go
*  创 建 者: mongia
*  创建日期：2021年12月28日
*
================================================================*/

package usersrv

import (
	"context"

	userpb "pharmacyerp/pb/user"
	usertask "pharmacyerp/task/user"
	"pharmacyerp/util"
)

// UserServer interface
type UserServer struct {
	userpb.UnimplementedUserServiceServer
}

// GetUserInfo 获取用户信息
func (*UserServer) GetUserInfo(ctx context.Context, req *userpb.GetUserInfoReq) (*userpb.GetUserInfoRes, error) {
	task := &usertask.GetUserInfoTask{
		Req: req,
		Res: &userpb.GetUserInfoRes{
			Data: &userpb.UserInfos{},
		},
		Err: util.EOK,
	}

	task.Run(ctx)
	return task.Res, nil
}

// GetUserInfoByID
func (*UserServer) GetUserInfoByID(ctx context.Context, req *userpb.GetUserInfoByIDReq) (*userpb.GetUserInfoByIDRes, error) {
	task := &usertask.GetUserInfoByIDTask{
		Req: req,
		Res: &userpb.GetUserInfoByIDRes{
			Data: &userpb.UserInfoDB{},
		},
		Err: util.EOK,
	}

	task.Run(ctx)
	return task.Res, nil
}

// DeleteUserInfo 删除用户
func (*UserServer) DeleteUserInfo(ctx context.Context, req *userpb.DeleteUserInfoReq) (*userpb.DeleteUserInfoRes, error) {
	task := &usertask.DeleteUserInfoTask{
		Req: req,
		Res: &userpb.DeleteUserInfoRes{},
		Err: util.EOK,
	}

	task.Run(ctx)
	return task.Res, nil
}

// CreateUserInfo 创建用户
func (*UserServer) CreateUserInfo(ctx context.Context, req *userpb.CreateUserInfoReq) (*userpb.CreateUserInfoRes, error) {
	task := &usertask.CreateUserInfoTask{
		Req: req,
		Res: &userpb.CreateUserInfoRes{
			Data: &userpb.UserInfoDB{},
		},
		Err: util.EOK,
	}

	task.Run(ctx)
	return task.Res, nil
}

// ModifyUserInfo 修改用户信息
func (*UserServer) ModifyUserInfo(ctx context.Context, req *userpb.ModifyUserInfoReq) (*userpb.ModifyUserInfoRes, error) {
	task := &usertask.ModifyUserInfoTask{
		Req: req,
		Res: &userpb.ModifyUserInfoRes{
			Data: &userpb.UserInfoDB{},
		},
		Err: util.EOK,
	}

	task.Run(ctx)
	return task.Res, nil
}
