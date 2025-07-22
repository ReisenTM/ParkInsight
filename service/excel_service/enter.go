package excel_service

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/xuri/excelize/v2"
	"parkAnalysis/global"
	"parkAnalysis/models"
	"strings"
	"sync"
	"time"
)

const maxConcurrency = 4

func ExcelUpdateHandler(filePath string) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	sheet := f.GetSheetName(0)
	rows, err := f.GetRows(sheet)
	if err != nil {
		panic(err)
	}
	//等待goroutine结束
	var wg sync.WaitGroup
	sem := make(chan struct{}, maxConcurrency)

	for i, row := range rows {
		if i == 0 || len(row) < 1 {
			continue
		}

		name := strings.TrimSpace(row[0])
		if name == "" {
			continue
		}

		rowIndex := i + 1 // Excel 行号从 1 开始

		wg.Add(1)
		sem <- struct{}{} // 占用一个并发槽

		go func(name string, rowIndex int) {
			defer wg.Done()
			defer func() { <-sem }() // 释放并发槽
			parkInfo := models.ParkModel{
				ParkName: name,
			}
			err = global.DB.Create(&parkInfo).Error
			if err != nil {
				logrus.Errorf("create parkInfo error:%v", err)
			} else {
				fmt.Printf("<=C 正在处理第%d个数据\n", i)
			}
			// 间隔,防止请求超时
			time.Sleep(500 * time.Millisecond)
		}(name, rowIndex)
	}

	wg.Wait()
	logrus.Println("全部完成")
}
