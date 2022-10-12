package zap

import (
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(path string, debug bool) *zap.Logger {
	encoder := getEncoder()

	var cores []zapcore.Core

	writeSyncer := getLogWriter(path)
	fileCore := zapcore.NewCore(encoder, writeSyncer, zapcore.InfoLevel)
	cores = append(cores, fileCore)

	if debug {
		consoleDebug := zapcore.Lock(os.Stdout)
		consoleCore := zapcore.NewCore(encoder, consoleDebug, zapcore.InfoLevel)
		cores = append(cores, consoleCore)
	}

	core := zapcore.NewTee(cores...)
	logger := zap.
		New(
			core,
			// zap.AddCaller(),
			// zap.AddCallerSkip(1),
		)
	return logger
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		// EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		// 	enc.AppendString(t.Format("2006-01-02 15:04:05"))
		// },
		// TimeKey:      "time",
		LevelKey: "level",
		// NameKey:      "logger",
		// CallerKey:    "caller",
		// MessageKey:   "msg",
		EncodeLevel:  zapcore.LowercaseLevelEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
	}
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(path string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   path + "/log.log",
		MaxSize:    2,
		MaxBackups: 10000,
		MaxAge:     0,
		Compress:   true,
	}
	return zapcore.AddSync(lumberJackLogger)
}
