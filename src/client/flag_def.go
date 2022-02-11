package main

import "flag"

var (
	address = flag.String("address", "localhost:9282", "grpc server address, localhost:9282")
	method  = flag.String("method", "", "rpc method: GetUserInfo...")
	param   = flag.String("param", "", "grpc request param in Json")
	timeout = flag.Int("timeout", 1, "rpc timeout in seconds")
)
