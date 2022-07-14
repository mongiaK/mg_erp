#!/bin/bash

#=================================================================
#  
#  文件名称：gen_client.sh
#  创 建 者: mongia
#  创建日期：2022年07月11日
#  邮    箱：mongiaK@outlook.com
#  
#=================================================================

go build -o user_cli main.go flag_def.go user_cli.go
