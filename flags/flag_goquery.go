package flags

import (
	"fmt"
	"parkAnalysis/service/spider_service"
	"time"
)

const baseUrl = "https://y.qianzhan.com/yuanqu/"

var ReqUrl = baseUrl

const pageNum = 77403 / 10 // 同时最多goroutine 工作
// FlagQuery 爬取园区内容
func FlagQuery() {
	ch := make(chan bool)

	for i := 1; i < pageNum; i++ {
		if i > 1 {
			ReqUrl = fmt.Sprintf("%s/?pg=%d", baseUrl, i)
		}

		go func(i int, url string) {

			spider_service.Spider(url, ch)
		}(i, ReqUrl)
		time.Sleep(100 * time.Millisecond)
	}
	for i := 1; i < pageNum; i++ {
		<-ch
	}
	fmt.Println("全部爬取成功")
}
