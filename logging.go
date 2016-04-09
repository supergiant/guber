package guber

import (
	"github.com/Sirupsen/logrus"
)

type logger struct {
	*logrus.Logger
}

func (l *logger) SetLevel(level string) {
	levelConst, err := logrus.ParseLevel(level)
	if err != nil {
		panic(err)
	}
	l.Level = levelConst
}

var Log *logger

func init() {
	Log = &logger{logrus.New()}
}
