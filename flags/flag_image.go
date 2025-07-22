package flags

import (
	"github.com/sirupsen/logrus"
	"parkAnalysis/global"
	"parkAnalysis/models"
	"parkAnalysis/service/image_service"
	"strings"
	"sync"
	"time"
)

const maxConcurrency = 8

func FlagImage() {
	var parks []models.ParkModel
	err := global.DB.Model(models.ParkModel{}).Find(&parks).Error
	if err != nil {
		logrus.Error(err)
		return
	}
	//等待goroutine结束
	var wg sync.WaitGroup
	sem := make(chan struct{}, maxConcurrency)
	for _, park := range parks {

		name := strings.TrimSpace(park.ParkName)
		if name == "" {
			continue
		}
		wg.Add(1)
		sem <- struct{}{} // 占用一个并发槽
		go func(name string) {
			defer wg.Done()
			defer func() { <-sem }() // 释放并发槽
			_ = image_service.DownloadFirstBingImage(name)

			// 间隔,防止请求超时
			time.Sleep(500 * time.Millisecond)
		}(name)
	}

	wg.Wait()
	logrus.Println("全部完成")

}
