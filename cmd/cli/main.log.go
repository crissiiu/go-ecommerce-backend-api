package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// //Sugar
	// sugar := zap.NewExample().Sugar()
	// sugar.Infof("Hello name:%s, age: %d", "thanh Sieu", 19)

	// //Logger
	// logger := zap.NewExample()
	// logger.Info("Hello", zap.String("name", "Thanh Sieu"), zap.Int("age", 20))

	// logger := zap.NewExample()

	// logger.Info("XinChao")
	// //Develpment
	// logger, _ = zap.NewDevelopment()
	// logger.Info("Hello Development")

	// //Production
	// logger, _ = zap.NewProduction()
	// logger.Info("Hello Production")

	encoder := getEncoderLog()
	sync := getWriterSync()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller())

	logger.Info("Info log", zap.Int("line", 1))
	logger.Error("Error log", zap.Int("line", 2))
}

// format log a message
func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()

	// 1791712152.6770558 --> 2026-10-11T16:49:12.676+0700
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	//ts --> Time
	encodeConfig.TimeKey = "time"

	//from info INFO
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	//"caller":"cli/main.log.go:22"
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encodeConfig)
}

func getWriterSync() zapcore.WriteSyncer {
	file, _ := os.OpenFile("./log/log.txt", os.O_CREATE|os.O_WRONLY, os.ModePerm)

	syncFile := zapcore.AddSync(file)

	syncConsole := zapcore.AddSync(os.Stderr)

	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}
