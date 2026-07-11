package pkg

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	loggerDevMode  = "dev"
	loggerProdMode = "prod"
)

func New(env string) *zap.Logger {
	var err error
	var logger *zap.Logger

	switch env {
	case loggerDevMode:
		config := zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		logger, err = config.Build()
	case loggerProdMode:
		config := zap.NewProductionConfig()
		logger, err = config.Build()
	default:
		log.Fatal("config dont have logger mode")
	}

	if err != nil {
		log.Fatal("error configuring logger: " + err.Error())
	}
	return logger
}
