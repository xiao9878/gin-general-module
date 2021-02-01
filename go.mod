module github.com/xiao9878/gin-general-module

go 1.15

require (
	github.com/casbin/casbin v1.9.1
	github.com/casbin/casbin/v2 v2.23.0
	github.com/casbin/gorm-adapter/v3 v3.1.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fsnotify/fsnotify v1.4.7
	github.com/gin-gonic/gin v1.6.3
	github.com/go-sql-driver/mysql v1.5.0
	github.com/satori/go.uuid v1.2.0
	github.com/spf13/viper v1.7.1
	go.uber.org/zap v1.16.0
	gorm.io/driver/mysql v1.0.3
	gorm.io/gorm v1.20.11
)

replace gorm.io/gorm v1.20.11 => ../go/pkg/mod/hzl.com/hzlorm
