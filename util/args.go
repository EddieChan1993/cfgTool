package util

import (
	"fmt"
	"os"
	"strings"
)

var ExportPath = "." //导出路径
var ExcelPath = "."  //配置表路径

func InitArgs() {
	if len(os.Args) < 3 {
		panic("args error")
	}
	ExcelPath = os.Args[1]
	ExportPath = os.Args[2]
	fmt.Println("ExcelPath", ExcelPath)
	fmt.Println("ExportPath", ExportPath)
}

func PkgName() string {
	slicePath := strings.Split(ExportPath, "/")
	return slicePath[len(slicePath)-1]
}
