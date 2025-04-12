package logger

import (
	"os"

	"github.com/alanbui/train-ticket/package/setting"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// wrap zap.Logger (now with Lumberjack) inside
type LoggerZap struct {
	*zap.Logger
}

// create a custom zap.Logger named LoggerZap using Lumberjack
func NewLogger(config setting.LoggerSetting) *LoggerZap {
	encoder := getEncoderLog() // defines how the log is formatted

	logLevel := config.Log_level
	var level zapcore.Level

	switch logLevel {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	default:
		level = zapcore.InfoLevel
	}

	// loads the logger config
	hook := lumberjack.Logger{
		Filename:   config.File_log_name,
		MaxSize:    config.Max_size,
		MaxBackups: config.Max_backups,
		MaxAge:     config.Max_age,  // days
		Compress:   config.Compress, // disabled by default
	}

	// defines the actual logging logic using Lumberjack.Logger
	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		level)

	// wraps the core in higher-level logger.
	// zapcore.Core doesnt have Info() or Error() so we need zap.Logger
	// logger := zap.New(core, zap.AddCaller()) // add caller to log caller
	logger := zap.New(core, zap.AddCaller())
	return &LoggerZap{logger}
}

// {"level":"info","ts":1743990553.6186502,"caller":"cli/main.log.go:8","msg":"Hello Production"}
// Format log
func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()

	encodeConfig.TimeKey = "time"

	// info => INFO
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// 1743990553.6186502 => 2025-04-06T16:16:04.877+0700
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// "caller":"cli/main.log.go:8"
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encodeConfig)
}
