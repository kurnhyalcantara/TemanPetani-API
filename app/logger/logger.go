package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

type Logrus struct {
	*logrus.Logger
}

func SetUpLogger() *Logrus {
	logger := &Logrus{logrus.New()}
	logger.Formatter = &logrus.TextFormatter{}
	logger.SetOutput(os.Stdout)

	return logger
}