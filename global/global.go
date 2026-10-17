package global

import (
	"github.com/crisiiu/go-ecommerce-backend-api/pkg/logger"
	"github.com/crisiiu/go-ecommerce-backend-api/pkg/settings"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config settings.Config
	Logger *logger.LoggerZap
	Rdb    *redis.Client
	Mdb    *gorm.DB
)

/*
	Config
	Redis
	MySql
	...
*/
