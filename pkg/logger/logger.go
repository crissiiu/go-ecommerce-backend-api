package logger

import (
	"os"

	"github.com/crisiiu/go-ecommerce-backend-api/pkg/settings"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerZap struct {
	*zap.Logger
}

func NewLogger(config settings.LoggerSetting) *LoggerZap {
	logLevel := config.Log_level
	var level zapcore.Level

	//debug -> info -> warning ->error ->fatal ->panic
	switch logLevel {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	case "fatal":
		level = zapcore.FatalLevel
	case "panic":
		level = zapcore.PanicLevel
	default:
		level = zapcore.InfoLevel
	}

	encoder := getEncoderLog()
	hook := lumberjack.Logger{
		Filename:   config.File_log_name,
		MaxSize:    config.Max_size,
		MaxBackups: config.Max_backups,
		MaxAge:     config.Max_age,
		Compress:   config.Compress,
	}
	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		level,
	)

	return &LoggerZap{zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))}
}

// format log
func getEncoderLog() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()

	// Change milisecond to date format
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// Change ts to Time
	encoderConfig.TimeKey = "time"

	// Information: INFO, ERROR, WARNING
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// "caller": "cli/main.log.go"
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encoderConfig)
}
