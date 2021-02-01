package global

import (
	"github.com/spf13/viper"
	"github.com/xiao9878/gin-general-module/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GVA_CONFIG config.Server
	GVA_DB     *gorm.DB
	GVA_VP     *viper.Viper
	GVA_ZAP    *zap.Logger
	GVA_LOG    *zap.Logger
)
