package initialize

import (
	"github.com/crisiiu/go-ecommerce-backend-api/global"
	"github.com/crisiiu/go-ecommerce-backend-api/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
