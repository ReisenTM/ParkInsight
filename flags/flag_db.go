package flags

import (
	"github.com/sirupsen/logrus"
	"parkAnalysis/global"
	"parkAnalysis/models"
)

// FlagDB 迁移数据库
func FlagDB() {

	err := global.DB.AutoMigrate(
		&models.ParkModel{},
		&models.ImageModel{})
	if err != nil {
		logrus.Errorf("自动迁移失败 %s", err)
		return
	}
	logrus.Infof("自动迁移成功")
}
