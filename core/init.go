package core

import (
	"cfgTool/excel"
	"cfgTool/util"
)

func initCore() {
	regIExcelThink(util.CfgConst, &excel.CfgConst{})
}
