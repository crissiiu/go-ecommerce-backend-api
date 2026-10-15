package global

import (
	"github.com/crisiiu/go-ecommerce-backend-api/pkg/logger"
	"github.com/crisiiu/go-ecommerce-backend-api/pkg/settings"
	"gorm.io/gorm"
)

var (
	Config settings.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
)

/*
	Config
	Redis
	MySql
	...
*/
