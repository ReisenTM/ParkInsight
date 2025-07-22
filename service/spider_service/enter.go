package spider_service

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"parkAnalysis/global"
	"parkAnalysis/models"
	"strings"
)

func Spider(url string, ch chan bool) {
	//1.发送请求
	client := http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/135.0.0.0 Safari/537.36")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9")
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//2.解析网页
	DetailDoc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		panic(err)
	}
	//3.获取节点信息
	//body > div.wrap > div.container > div.box-s2.mt15 > table > tbody > tr:nth-child(2) > td.f-tal > a
	//body > div.wrap > div.container > div.box-s2.mt15 > table > tbody > tr:nth-child(1)
	//body > div.wrap > div.container > div.box-s2.mt15 > table > tbody > tr:nth-child(2) > td.f-tal > a
	//body > div.wrap > div.container > div.box-s2.mt15 > table > tbody > tr:nth-child(1)
	//body > div.wrap > div.container > div.box-s2.mt15 > table > tbody > tr:nth-child(1) > td.f-tal
	DetailDoc.Find("body > div.wrap > div.container > div.box-s2.mt15> table > tbody").
		Each(func(i int, s *goquery.Selection) {
			//body > div.wrap > div.container > div.box-s2.mt15 > table > tbody > tr:nth-child(1) > td.f-tal > a
			data := s.Find("a").Text()
			//4.保存内容

			res := strings.Split(data, "查看")
			for i, name := range res {
				if name != "" {
					temp := name
					fmt.Printf("第%d个是%s\n", i, temp)
					model := models.ParkModel{
						ParkName: temp,
					}
					//先查存在不存在
					// 保存到数据库
					if err := global.DB.Create(&model).Error; err != nil {
						fmt.Printf("保存第 %d 个元素失败: %v\n", i, err)
					} else {
						fmt.Printf("保存第 %d 个元素成功\n", i)
					}
				}
			}
		})
	if ch != nil {
		ch <- true
	}
}
