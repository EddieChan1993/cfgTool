package excel

import (
	"cfgTool/util"
	"fmt"
	"github.com/xuri/excelize/v2"
	"os"
	"strings"
)

type CfgConst struct {
}

func (c *CfgConst) Think(excelFileName string, f *excelize.File) (goFileName string, err error) {
	rows, err := f.GetRows("const")
	if err != nil {
		fmt.Println(err)
		return
	}
	goFileName = fmt.Sprintf("%sSuper.go", strings.ToLower(excelFileName))
	OSFile, err := os.OpenFile(fmt.Sprintf("%s/%s", util.ExportPath, goFileName), os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return
	}
	defer OSFile.Close()
	//覆盖写入
	OSFile.Truncate(0)
	n, _ := OSFile.Seek(0, 0)
	//文件头信息
	pkgName := ""
	if util.IsNoBuild != "true" {
		pkgName = fmt.Sprintf("package %s\n", util.PkgName())
	} else {
		pkgName = fmt.Sprintf("//go:build !linux && !darwin && !windows\npackage %s\n", util.PkgName())
	}
	_, err = OSFile.WriteAt([]byte(pkgName), n)
	n += int64(len([]byte(pkgName)))
	//const变量声明
	_, err = OSFile.WriteAt([]byte("const (\n"), n)
	n += int64(len([]byte("const (\n")))
	var msg string
	for i, row := range rows[5:] {
		if len(row) <= 3 {
			panic(fmt.Sprintf("row %d format error %v", i+6, row))
		}
		if row[3] == "" {
			msg = fmt.Sprintf("CST_%s=%s //%s %s\n", strings.ToUpper(strings.Replace(row[1], " ", "X", 1)), row[0], row[2], row[4])
		} else {
			msg = fmt.Sprintf("CST_%s=%s //%s %s\n", strings.ToUpper(strings.Replace(row[1], " ", "X", 1)), row[0], row[2], row[3])
		}
		_, err = OSFile.WriteAt([]byte(msg), n)
		if err != nil {
			panic(fmt.Sprintf("err %v msg %v", err, msg))
		}
		n += int64(len([]byte(msg)))
	}
	_, err = OSFile.WriteAt([]byte(")"), n)
	return goFileName, nil
}
