package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path/filepath"
)

func New() func() {
	var logger *zap.Logger
	switch os.Getenv("RUN_MODE") {
	case "local":
		atom := zap.NewAtomicLevel()
		encoderCfg := zap.NewProductionEncoderConfig()
		encoderCfg.TimeKey = "timestamp"
		encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
		encoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
		logger = zap.New(zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderCfg),
			zapcore.Lock(os.Stdout),
			atom,
		), zap.AddCaller())
	default:
		logger, _ = zap.NewProduction()
	}
	return zap.ReplaceGlobals(logger)
}

func NewFileLogger(path, appname string) func() {

	filename := filepath.Join(path, appname, "app.log")
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   filename,
		MaxSize:    50,
		MaxBackups: 90,
		MaxAge:     28,
		Compress:   true,
	})
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		w,
		zap.InfoLevel,
	)
	logger := zap.New(core)
	return zap.ReplaceGlobals(logger)
}
