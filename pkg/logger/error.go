package logger

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func (logger *Log) WithErr(err error) *logrus.Entry {
	switch x := err.(type) {
	case error:
		fmt.Errorf("error print", x)
		return logrus.NewEntry(logger.Logger)
	default:
		return logger.Logger.WithError(err)
	}
}


