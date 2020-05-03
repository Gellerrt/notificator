package notificator

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"notificator/internal/config"
)

type Notificator struct {
	version string
	conf    *config.Config
	logger  *logrus.Logger
}

func New(c *config.Config) *Notificator {
	return &Notificator{
		version: "0.0.1",
		conf:    c,
		logger:  logrus.New(),
	}
}

func (n *Notificator) Start() error {
	n.logger.Info(fmt.Sprintf("<-------------------- Application version %s is starting -------------------->", n.version))
	if err := n.configLogger(); err != nil {
		return err
	}
	n.logger.Info(fmt.Sprintf("<-------------------- Application started -------------------->"))
	return nil
}

func (n *Notificator) configLogger() error {
	level, err := logrus.ParseLevel(n.conf.LogLevel)
	if err != nil {
		return err
	}
	n.logger.SetLevel(level)
	n.logger.AddHook()
	return nil
}
