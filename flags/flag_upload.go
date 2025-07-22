package flags

import (
	"parkAnalysis/service/excel_service"
	"strings"
)

func FLagUpload(filePath string) {
	if strings.HasSuffix(filePath, ".xlsx") {
		excel_service.ExcelUpdateHandler(filePath)
	}
	if strings.HasSuffix(filePath, ".csv") {
		//TODO：CSV上传
	}
}
