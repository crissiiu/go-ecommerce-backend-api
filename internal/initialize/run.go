package initialize

import (
	"fmt"

	"github.com/crisiiu/go-ecommerce-backend-api/global"
)

func Run() {
	//load configuration
	LoadConfig()
	fmt.Println("Load Config MySQL", global.Config.MySql.Username)
	InitLogger()
	InitMySql()
	InitRouter()
	r := InitRouter()
	r.Run(":8002")
}
