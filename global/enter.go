package global

import (
	"gorm.io/gorm"
	"parkAnalysis/conf"
)

var (
	Config *conf.Config //全局配置
	DB     *gorm.DB     //数据库
)
