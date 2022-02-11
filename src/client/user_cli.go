/*================================================================
*
*  文件名称：user_cli.go
*  创 建 者: mongia
*  创建日期：2021年03月04日
*
================================================================*/

package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "pharmacyerp/pb/user"
)

func init() {
	registerFunc("userCreate", userCreate)
	registerFunc("userGet", userGet)
	registerFunc("userGetOne", userGetOne)
	registerFunc("userDelete", userDelete)
	registerFunc("userModify", userModify)
}

func userModify(conn grpc.ClientConnInterface, data []byte) {
	c := pb.NewUserServiceClient(conn)

	req := &pb.ModifyUserInfoReq{}

	err := json.Unmarshal(data, req)
	if nil != err {
		log.Fatalf("input param is not json or format error, err: %+v", err)
	}

	log.Printf("request is : %+v", req)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	r, err := c.ModifyUserInfo(ctx, req)
	if err != nil {
		log.Printf("could not rpc request: %v", err)
		return
	}
	log.Printf("rpc result: %+v", r)
}

func userCreate(conn grpc.ClientConnInterface, data []byte) {
	c := pb.NewUserServiceClient(conn)

	req := &pb.CreateUserInfoReq{}

	err := json.Unmarshal(data, req)
	if nil != err {
		log.Fatalf("input param is not json or format error, err: %+v", err)
	}

	log.Printf("request is : %+v", req)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	r, err := c.CreateUserInfo(ctx, req)
	if err != nil {
		log.Printf("could not rpc request: %v", err)
		return
	}
	log.Printf("rpc result: %+v", r)
}

func userDelete(conn grpc.ClientConnInterface, data []byte) {
	c := pb.NewUserServiceClient(conn)

	req := &pb.DeleteUserInfoReq{}
	err := json.Unmarshal(data, req)
	if nil != err {
		log.Fatalf("input param is not json or format error, err: %+v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	r, err := c.DeleteUserInfo(ctx, req)
	if err != nil {
		log.Printf("could not rpc request: %v", err)
	}
	log.Printf("rpc result: %+v", r)
}

func userGetOne(conn grpc.ClientConnInterface, data []byte) {
	c := pb.NewUserServiceClient(conn)

	req := &pb.GetUserInfoByIDReq{}
	err := json.Unmarshal(data, req)
	if nil != err {
		log.Fatalf("input param is not json or format error, err: %+v", err)
	}

	log.Printf("request is : %+v", req)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	r, err := c.GetUserInfoByID(ctx, req)
	if err != nil {
		log.Printf("could not rpc request: %v", err)
		return
	}
	log.Printf("rpc result: %+v", r)
}

func userGet(conn grpc.ClientConnInterface, data []byte) {
	c := pb.NewUserServiceClient(conn)

	req := &pb.GetUserInfoReq{}
	err := json.Unmarshal(data, req)
	if nil != err {
		log.Fatalf("input param is not json or format error, err: %+v", err)
	}

	log.Printf("request is : %+v", req)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	r, err := c.GetUserInfo(ctx, req)
	if err != nil {
		log.Printf("could not rpc request: %v", err)
		return
	}
	log.Printf("rpc result: %+v", r)
}
