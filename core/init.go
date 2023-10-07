package core

import (
	"cfgTool/excel"
	"cfgTool/util"
)

func initCore() {
	RegIExcelThink(util.CfgConst, &excel.CfgConst{})
}
