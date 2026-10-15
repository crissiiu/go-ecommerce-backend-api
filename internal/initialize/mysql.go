package initialize

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"

	"github.com/crisiiu/go-ecommerce-backend-api/global"
	"github.com/crisiiu/go-ecommerce-backend-api/internal/po"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func checkErrorPanic(err error, errStr string) {
	if err != nil {
		global.Logger.Error(errStr, zap.Error(err))
		panic(err)
	}
}

func InitMySql() {
	m := global.Config.MySql
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.DbName)
	db, err := gorm.Open(mysql.Open(s), &gorm.
		Config{
		SkipDefaultTransaction: true,
	})
	checkErrorPanic(err, "InitMySql initialization failed")
	global.Logger.Info("Initialazing MySql Successfully")
	global.Mdb = db

	//Set pool
	SetPool()

	migrateTable()
}

func SetPool() {
	m := global.Config.MySql

	sqlDB, err := global.Mdb.DB()

	if err != nil {
		fmt.Println("MySql Error: ", err)
	}

	sqlDB.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))

	sqlDB.SetMaxOpenConns(m.MaxOpenConns)

	sqlDB.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime))

}

func migrateTable() {
	err := global.Mdb.AutoMigrate(
		&po.User{},
		&po.Role{},
	)
	if err != nil {
		fmt.Println("MySql Migration Table Error: ", err)
	}
}
