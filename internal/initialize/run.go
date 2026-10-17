package initialize

import (
	"fmt"

	"github.com/crisiiu/go-ecommerce-backend-api/global"
	"go.uber.org/zap"
)

func Run() {
	//load configuration
	LoadConfig()
	fmt.Println("Load Config MySQL", global.Config.MySql.Username)
	InitLogger()
	global.Logger.Info("Config log ok", zap.String("ok", "success"))
	InitMySql()
	InitRedis()
	InitRouter()
	r := InitRouter()
	r.Run(":8002")
}
