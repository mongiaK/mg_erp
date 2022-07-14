/*================================================================
*
*  文件名称：gen_rpc_proto.go
*  创 建 者: mongia
*  创建日期：2022年07月14日
*  邮    箱：mongiaK@outlook.com
*
================================================================*/

package main

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
)

var (
	serviceName = ""
	moduleName  = ""
	projectName = ""
)

func generateGrpcServerProtoFile(gen *protogen.Plugin, file *protogen.File) {
	filename := projectName + "/proto/" + strings.ToLower(moduleName) + "_service.proto"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	g.P("syntax=\"proto3\";")
	g.P()
	g.P("import \"", file.GoPackageName, "_msg.proto\";")
	g.P()
	g.P("option go_package=", file.GoImportPath, ";")
	g.P()

	g.P("service ", serviceName, " {")
	g.P("	rpc Get", moduleName, "(Get", moduleName, "Req) returns(Get", moduleName, "Res) {};")
	g.P("	rpc Delete", moduleName, "(Delete", moduleName, "Req) returns(Delete", moduleName, "Res) {};")
	g.P("	rpc Modify", moduleName, "(Modify", moduleName, "Req) returns(Modify", moduleName, "Res) {};")
	g.P("	rpc Create", moduleName, "(Create", moduleName, "Req) returns(Create", moduleName, "Res) {};")
	g.P("}")
}

func generateGrpcMsgProtoFile(gen *protogen.Plugin, file *protogen.File) {
	filename := projectName + "/proto/" + strings.ToLower(moduleName) + "_msg.proto"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	g.P("syntax=\"proto3\";")
	g.P()
	g.P("option go_package=", file.GoImportPath, ";")
	g.P()

	g.P("enum Logic {")
	g.P("	LESS = 0;")
	g.P("	LESS_THAN = 1;")
	g.P("	EQUAL = 2;")
	g.P("	GREAT_THAN = 3;")
	g.P("	GREAT = 4;")
	g.P("	IN = 5;")
	g.P("	NOT_IN = 6;")
	g.P("	GROUP_BY = 7;")
	g.P("	ORDER_BY = 8;")
	g.P("	NOT_EQUAL = 9;")
	g.P("}")
	g.P()

	g.P("enum ", moduleName, "Item {")
	g.P("	ALL = 0;")
	for _, item := range file.Messages[0].Fields {
		g.P("	", strings.ToUpper(item.GoName), " = ", 1<<item.Desc.Number(), ";")
	}
	g.P("}")
	g.P()

	g.P("message ", moduleName, "DB {")
	for _, item := range file.Messages[0].Fields {
		g.P("	", item.Desc.Kind().String(), " ", item.Desc.Name(), " = ", item.Desc.Number(), ";")
	}
	g.P("}")
	g.P()

	g.P("enum ConditionValueType {")
	g.P("	NULL = 0;")
	g.P("	INT = 1;")
	g.P("	STRING = 2;")
	g.P("	INT_ARRAY = 3;")
	g.P("	STRING_ARRAY = 4;")
	g.P("}")
	g.P()

	g.P("message ConditionItem {")
	g.P("	", moduleName, "Item key = 1;")
	g.P("	Logic logic = 2;")
	g.P("	ConditionValueType vtype = 3;")
	g.P("	int64 ivalue = 10;")
	g.P("	string svalue = 11;")
	g.P("	repeated int64 iavalue = 12;")
	g.P("	repeated string savalue = 13;")
	g.P("}")
	g.P()

	g.P("message ConditionItemArray {")
	g.P("	repeated ConditionItem items = 1;")
	g.P("}")
	g.P()

	g.P("message Condition {")
	g.P("	ConditionItemArray and = 1;")
	g.P("	ConditionItemArray or = 2;")
	g.P("	ConditionItem orderby = 3;")
	g.P("	ConditionItem groupby = 4;")
	g.P("	int32 pagesize = 5;")
	g.P("	int32 pagenum = 6;")
	g.P("}")
	g.P()

	g.P("message ", moduleName, "s {")
	g.P("	int64 count = 1;")
	g.P("	repeated ", moduleName, "DB ", strings.ToLower(moduleName), " = 2;")
	g.P("}")
	g.P()

	generateGetMsg(g)
	generateCreateMsg(g)
	generateDeleteMsg(g)
	generateModifyMsg(g)
}

func generateGetMsg(g *protogen.GeneratedFile) {
	reqMsgName := "Get" + moduleName + "Req"
	resMsgName := "Get" + moduleName + "Res"

	g.P("message ", reqMsgName, " {")
	g.P("	int64 items = 1;")
	g.P("	Condition cond = 2;")
	g.P("}")
	g.P()

	g.P("message ", resMsgName, " {")
	g.P("	int32 code = 1;")
	g.P("	string msg = 2;")
	g.P("	", moduleName, "s data = 3;")
	g.P("}")
	g.P()
}

func generateDeleteMsg(g *protogen.GeneratedFile) {
	reqMsgName := "Delete" + moduleName + "Req"
	resMsgName := "Delete" + moduleName + "Res"

	g.P("message ", reqMsgName, " {")
	g.P("	Condition cond = 1;")
	g.P("}")
	g.P()

	g.P("message ", resMsgName, " {")
	g.P("	int32 code = 1;")
	g.P("	string msg = 2;")
	g.P("}")
	g.P()
}

func generateModifyMsg(g *protogen.GeneratedFile) {
	reqMsgName := "Modify" + moduleName + "Req"
	resMsgName := "Modify" + moduleName + "Res"

	g.P("message ", reqMsgName, " {")
	g.P("	Condition cond = 1;")
	g.P("	", moduleName, "DB  data = 2;")
	g.P("	int64 items = 3;")
	g.P("}")
	g.P()

	g.P("message ", resMsgName, " {")
	g.P("	int32 code = 1;")
	g.P("	string msg = 2;")
	g.P("}")
	g.P()
}

func generateCreateMsg(g *protogen.GeneratedFile) {
	reqMsgName := "Create" + moduleName + "Req"
	resMsgName := "Create" + moduleName + "Res"

	g.P("message ", reqMsgName, " {")
	g.P("	", moduleName, "DB ", strings.ToLower(moduleName), " = 1;")
	g.P("}")
	g.P()

	g.P("message ", resMsgName, " {")
	g.P("	int32 code = 1;")
	g.P("	string msg = 2;")
	g.P("	", moduleName, "DB data = 3;")
	g.P("}")
	g.P()
}

func generateRPCProtoFile(gen *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile {
	if len(file.Messages) != 1 {
		fmt.Printf("just need one message segment\n")
		return nil
	}

	moduleName = file.Messages[0].GoIdent.GoName
	serviceName = file.Messages[0].GoIdent.GoName + "Service"
	projectName = strings.ToLower(moduleName)

	generateGrpcServerProtoFile(gen, file)
	generateGrpcMsgProtoFile(gen, file)
	return nil
}
