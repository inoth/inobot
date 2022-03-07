package logger

import "github.com/sirupsen/logrus"

type LogrusConfig struct {
	Hooks []logrus.Hook
}

func (l *LogrusConfig) AddHook(hook ...logrus.Hook) {
	l.Hooks = hook
}

func (l *LogrusConfig) Init() error {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	for _, hook := range l.Hooks {
		logrus.AddHook(hook)
	}
	return nil
}
