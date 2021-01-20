package global

import (
	"gin-general-module/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GVA_CONFIG config.Server
	GVA_DB     *gorm.DB
	GVA_VP     *viper.Viper
	GVA_ZAP    *zap.Logger
)
