package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

// InitLogger 初始化日志工具
func InitLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logger.SetLevel(logrus.InfoLevel)
	return logger
}
