/*================================================================
*
*  文件名称：main.go
*  创 建 者: mongia
*  创建日期：2022年07月01日
*  邮    箱：mongiaK@outlook.com
*
================================================================*/

package main

import (
	"flag"
	"fmt"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

const version = "1.0.0"

var requireUnimplemented *bool

// FirstUpper 字符串首字母大写
func firstUpper(s string) string {
	if s == "" {
		return ""
	}

	return strings.ToUpper(s[:1]) + s[1:]
}

// FirstLower 字符串首字母小写
func firstLower(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}

func main() {
	showVersion := flag.Bool("version", false, "print the version and exit")
	language := flag.String("language", "go", "set language you need to generate, default: go")
	flag.Parse()
	if *showVersion {
		fmt.Printf("protoc-gen-curd %v\n", version)
		return
	}

	var flags flag.FlagSet
	requireUnimplemented = flags.Bool("require_unimplemented_servers", true, "set to false to match legacy behavior")

	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		if len(gen.Files) != 1 {
			fmt.Printf("input one file")
			return nil
		}

		f := gen.Files[0]
		if !f.Generate {
			fmt.Printf("go generate faild\n")
			return nil
		}

		generateRPCProtoFile(gen, f)

		switch *language {
		case "go":
			generateGoCurd(gen, f)
			break
		case "cpp":
		case "java":
		default:
			fmt.Printf("%s not surport\n", *language)
			return nil
		}
		return nil
	})

}
