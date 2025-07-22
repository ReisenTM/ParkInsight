package flags

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"parkAnalysis/global"
	"parkAnalysis/models"
	"parkAnalysis/service/ai_service"
	"strings"
	"sync"
)

type ParkInfo struct {
	ParkName      string `json:"parkName"`
	Level         string `json:"level"`           // 园区级别
	ParkType      string `json:"parkType"`        // 园区类型
	LandProperty  string `json:"natureOfLand"`    // 土地性质
	MainIndustry  string `json:"leadingIndustry"` // 主导产业
	Introduce     string `json:"introduce"`       // 园区介绍
	Advantage     string `json:"advantage"`       // 园区优势
	EstablishTime string `json:"establishTime"`   // 成立时间

}

const maxProcessor = 4

// FlagAnalysis 进行数据分析并填充
func FlagAnalysis() {
	var count int64
	global.DB.Model(models.ParkModel{}).Count(&count)
	var mark int64 = 1
	var parks []models.ParkModel
	err := global.DB.Model(models.ParkModel{}).Find(&parks).Error
	if err != nil {
		logrus.Errorf("find park err: %v", err)
	}
	wg := sync.WaitGroup{}
	sem := make(chan struct{}, maxProcessor)
	for _, park := range parks {
		fmt.Printf("<=C 正在处理第%d个数据\n", mark)
		mark++
		sem <- struct{}{} //占用一个槽
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer func() { <-sem }() //释放信号量
			var existing models.ParkModel
			err := global.DB.Model(&models.ParkModel{}).
				Where("park_name = ? AND (introduce IS NULL OR introduce = '')", park.ParkName).
				First(&existing).Error
			if err != nil {
				return // 已有内容，不处理
			}
			res, _ := ai_service.DSToChat(park.ParkName)
			msg1, _ := strings.CutPrefix(res, "```json")
			final, _ := strings.CutSuffix(msg1, "```")
			var parkInfo ParkInfo
			_ = json.Unmarshal([]byte(final), &parkInfo)
			updates := map[string]interface{}{
				"level":          parkInfo.Level,
				"park_type":      parkInfo.ParkType,
				"land_property":  parkInfo.LandProperty,
				"main_industry":  parkInfo.MainIndustry,
				"introduce":      parkInfo.Introduce,
				"advantage":      parkInfo.Advantage,
				"establish_time": parkInfo.EstablishTime,
			}
			if updates != nil {
				global.DB.Model(park).Updates(updates)
			}
		}()
	}

}
