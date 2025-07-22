package main

import (
	"parkAnalysis/core"
	"parkAnalysis/flags"
	"parkAnalysis/global"
	"parkAnalysis/router"
)

func main() {
	flags.Parse()                   //绑定命令行参数
	global.Config = core.ReadConf() //读配置文件
	core.InitDefaultLogus()         //初始化日志
	global.DB = core.InitDb()       //初始化数据库
	flags.Run()
	//启动程序

	router.Run()
}
