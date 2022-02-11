/*================================================================
*
*  文件名称：main.go
*  创 建 者: mongia
*  创建日期：2021年12月28日
*
================================================================*/

package main

import (
	"fmt"
	"net"

	"github.com/sevlyar/go-daemon"
	dm "github.com/sevlyar/go-daemon"
	"google.golang.org/grpc"

	usersrv "pharmacyerp/app/user"
	userpb "pharmacyerp/pb/user"

	config "pharmacyerp/config"
	log "pharmacyerp/log"
	third "pharmacyerp/third"
)

func prepare() {
	log.InitLog()

	err := config.LoadConfig()
	if nil != err {
		log.Fatal("load config failed",
			log.String("err", err.Error()))
	}

	err = third.MysqlInit(config.Cfg.Mysql.Database, config.Cfg.Mysql.Username, config.Cfg.Mysql.Password, config.Cfg.Mysql.Host, config.Cfg.Mysql.Port, config.Cfg.Mysql.Conns)
	if nil != err {
		log.Fatal("mysql init failed",
			log.String("err", err.Error()))
	}

	third.RedisInit(fmt.Sprintf("%s:%d", config.Cfg.Redis.Host, config.Cfg.Redis.Port), config.Cfg.Redis.Password, config.Cfg.Redis.Conns)
}

func daemonsize() *daemon.Context {
	cntxt := &dm.Context{
		PidFilePerm: 0644,
		WorkDir:     "runtime",
		Umask:       027,
	}

	d, err := cntxt.Reborn()
	if err != nil {
		log.Fatal("run daemon failed")
		return nil
	}
	if d != nil {
		log.Fatal("father run over")
		return nil
	}
	return cntxt
}

func main() {
	prepare()

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", config.Cfg.Host, config.Cfg.Port))

	if nil != err {
		log.Fatal("falied to listen", log.String("err", err.Error()))
	}

	grpcServer := grpc.NewServer()

	userpb.RegisterUserServiceServer(grpcServer, &usersrv.UserServer{})

	log.Info("server run")
	grpcServer.Serve(lis)
}
