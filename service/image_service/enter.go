package image_service

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func DownloadFirstBingImage(query string) error {
	fmt.Println("正在处理:", query)
	searchURL := "https://www.bing.com/images/search?q=" + url.QueryEscape(query)

	resp, err := http.Get(searchURL)
	if err != nil {
		return fmt.Errorf("请求失败: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP请求失败: %s", resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return fmt.Errorf("解析 HTML 失败: %v", err)
	}

	var imageURL string
	doc.Find("img").EachWithBreak(func(i int, s *goquery.Selection) bool {
		src, exists := s.Attr("src")
		if exists && strings.HasPrefix(src, "http") && !strings.HasSuffix(src, ".gif") {
			imageURL = src
			return false // 找到就跳出循环
		}
		return true
	})

	if imageURL == "" {
		return fmt.Errorf("未找到图片链接")
	}
	fmt.Println("找到图片：", imageURL)

	// 下载图片
	imageResp, err := http.Get(imageURL)
	if err != nil {
		return fmt.Errorf("下载图片失败: %v", err)
	}
	defer imageResp.Body.Close()

	if imageResp.StatusCode != http.StatusOK {
		return fmt.Errorf("图片下载失败: %s", imageResp.Status)
	}

	// 创建文件并保存
	dir, _ := os.Getwd()
	folder := fmt.Sprintf("%s/images/%s", dir, query)

	// 创建文件夹（如果不存在）
	err = os.MkdirAll(folder, 0755)
	if err != nil {
		log.Println("Create folder failed:", err)
		return err
	}
	fileName := fmt.Sprintf("%s/%s.jpg", folder, query)
	if fileExists(fileName) {
		//存在就跳过
		return nil
	}
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("创建文件失败: %v", err)
	}
	defer file.Close()

	_, err = io.Copy(file, imageResp.Body)
	if err != nil {
		return fmt.Errorf("保存图片失败: %v", err)
	}

	fmt.Println("保存成功:", fileName)
	return nil
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || !os.IsNotExist(err)
}
