package logger

import (
	"fmt"
	"os"

	"go.uber.org/zap"
)

type logger struct {
	*zap.Logger
}

// シングルトン
var singleton = newLogger()

func newLogger() *logger {
	env := os.Getenv("ENV")
	if env == "prod" {
		prod, _ := zap.NewProduction()
		return &logger{prod}
	}
	zap, _ := zap.NewDevelopment()
	return &logger{zap}
}

// Debugログ
func Debug(msg string) {
	env := os.Getenv("ENV")
	if env == "prod" {
		return
	}
	singleton.Debug(msg)
}

// Infoログ
func Info(msg string) {
	singleton.Info(msg)
}

// Warnログ
func Warn(msg string) {
	singleton.Warn(msg)
}

// Errorログ
func Error(msg string) {
	singleton.Error(msg)
}

func StartLog() {
	fmt.Println("grpc server start")
}
