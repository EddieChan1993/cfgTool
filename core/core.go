package core

import (
	"cfgTool/util"
	"fmt"
	"github.com/xuri/excelize/v2"
)

var AllExcelThinks = make(map[string]IExcelThink)

type IExcelThink interface {
	Think(excelFileName string, f *excelize.File) (goFileName string, err error)
}

func RegIExcelThink(regName string, obj IExcelThink) {
	AllExcelThinks[regName] = obj
}

func RunCore() {
	initCore()
	var err error
	var path string
	var fileName string
	for cfgName, iThink := range AllExcelThinks {
		var f *excelize.File
		path = fmt.Sprint(util.ExcelPath + "/" + cfgName + ".xlsx")
		f, err = excelize.OpenFile(path)
		if err != nil {
			fmt.Printf("OpenFile path %s err %v\n", path, err)
			return
		}
		defer func() {
			if err = f.Close(); err != nil {
				fmt.Println("Close", err)
			}
		}()
		fileName, err = iThink.Think(cfgName, f)
		if err != nil {
			fmt.Println("Think", err)
			return
		}
		err = util.Command(fmt.Sprintf("go fmt %s/%s", util.ExportPath, fileName))
		if err != nil {
			fmt.Println("Command", err)
			return
		}
		fmt.Println(path, " Ok")
	}
}
