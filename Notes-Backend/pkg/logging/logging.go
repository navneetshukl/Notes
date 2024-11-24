package logging

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Log struct {
	Log *zap.Logger
}

type Logger interface {
	Info(msg string)
	Error(msg string, err error)
}

func NewLogger() *Log {
	// Define the encoder configuration for both console and file outputs
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderCfg.CallerKey = "caller"

	consoleEncoder := zapcore.NewConsoleEncoder(encoderCfg)

	jsonEncoder := zapcore.NewJSONEncoder(encoderCfg)

	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return nil
	}

	consoleSyncer := zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zap.NewAtomicLevelAt(zapcore.InfoLevel))
	fileSyncer := zapcore.NewCore(jsonEncoder, zapcore.AddSync(file), zap.NewAtomicLevelAt(zapcore.InfoLevel))

	core := zapcore.NewTee(consoleSyncer, fileSyncer)

	logger := zap.New(core, zap.AddStacktrace(zapcore.ErrorLevel))

	return &Log{
		Log: logger,
	}
}

// Custom function to log the full file path and line number
func fullPathCallerEncoder() string {
	_, filename, line, ok := runtime.Caller(2) 
	if !ok {
		return "unknown-caller"
	}

	// Convert to absolute path
	absPath, err := filepath.Abs(filename)
	if err != nil {
		return filename
	}

	return fmt.Sprintf("%s:%d", filepath.ToSlash(absPath), line)
}

func (l *Log) Info(msg string) {
	l.Log.Info(msg, zap.String("caller", fullPathCallerEncoder()))
}

func (l *Log) Error(msg string, err error) {
	message := fmt.Sprintf("%s %v", msg, err)
	l.Log.Error(message, zap.String("caller", fullPathCallerEncoder()))
}
