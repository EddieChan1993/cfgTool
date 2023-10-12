package util

import (
	"fmt"
	"os"
	"strings"
)

var ExportPath = "."    //导出路径
var ExcelPath = "."     //配置表路径
var IsNoBuild = "false" //不参与编译

func InitArgs() {
	if len(os.Args) < 3 {
		panic("args error")
	}
	ExcelPath = os.Args[1]
	ExportPath = os.Args[2]
	lastArg := os.Args[len(os.Args)-1]
	if lastArg != ExportPath {
		IsNoBuild = lastArg
	}
	fmt.Println("ExcelPath", ExcelPath)
	fmt.Println("ExportPath", ExportPath)
	fmt.Println("IsNoBuild", IsNoBuild)
}

func PkgName() string {
	slicePath := strings.Split(ExportPath, "/")
	return slicePath[len(slicePath)-1]
}
