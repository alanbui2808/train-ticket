package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// LOG INFO
func main() {
	encoder := getEncoderLog() // defines how the log is formatted
	sync := getWriterSync()    // determines where the logs are written (e.g: giles, stderr)

	// defines the actual logging logic
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)

	// wraps the core in higher-level logger.
	// zapcore.Core doesnt have Info() or Error() so we need zap.Logger
	logger := zap.New(core, zap.AddCaller()) // add caller to log caller

	/*
		1. zapcore.Core:
			- Encode the logs (zapcore.Encoder)
			- Where to write them (zapcore.WriteSyncer)
			- What level of logs to allow (zapcore.Level)

		2. zapcore.Logger:
		 user-facing logging object that:
		 	- Provides methods like Info(), Error(), Debug() etc
			- Adds features like callinto info AddCaller(), stack traces, etc.
			- Internally delegates to the core for the actual work.
	*/
	logger.Info("Info log", zap.Int("line", 1))
	logger.Error("Info log", zap.Int("line", 2))
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

func getWriterSync() zapcore.WriteSyncer {
	file, _ := os.OpenFile("./log/log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)

	// wraps file in a WriteSyncer so zapcore can use it as input destination
	syncFile := zapcore.AddSync(file)         // file
	syncConsole := zapcore.AddSync(os.Stderr) // console

	// Combines both WriteSyncer
	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}
