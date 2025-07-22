package flags

import (
	"flag"
	"os"
)

type Options struct {
	File     string
	DB       bool
	Query    bool
	Image    bool
	Upload   string
	Analysis bool
}

var FlagOptions = new(Options)

// Parse flag绑定
func Parse() {
	flag.StringVar(&FlagOptions.File, "f", "settings.yaml", "配置文件")
	flag.BoolVar(&FlagOptions.DB, "db", false, "数据库迁移")
	flag.BoolVar(&FlagOptions.Query, "q", false, "爬取数据")
	flag.StringVar(&FlagOptions.Upload, "u", "", "上传数据")
	flag.BoolVar(&FlagOptions.Analysis, "a", false, "数据分析")
	flag.BoolVar(&FlagOptions.Image, "i", false, "图片下载")
	flag.Parse()
}

// Run flag实现
func Run() {
	if FlagOptions.DB {
		//执行数据库迁移
		FlagDB()
		os.Exit(0)
	}
	if FlagOptions.Query {
		//爬取产业园区信息
		FlagQuery()
		os.Exit(0)
	}
	if FlagOptions.Upload != "" {
		FLagUpload(FlagOptions.Upload)
		os.Exit(0)
	}
	if FlagOptions.Analysis {
		FlagAnalysis()
		os.Exit(0)
	}
	if FlagOptions.Image {
		FlagImage()
		os.Exit(0)
	}
}
