package global

import (
	"github.com/crisiiu/go-ecommerce-backend-api/pkg/logger"
	"github.com/crisiiu/go-ecommerce-backend-api/pkg/settings"
)

var (
	Config settings.Config
	Logger *logger.LoggerZap
)

/*
	Config
	Redis
	MySql
	...
*/
