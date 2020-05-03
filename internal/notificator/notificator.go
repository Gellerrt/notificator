package notificator

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"notificator/internal/store"
	"os"
	"time"
)

type Notificator struct {
	version string
	conf    *Config
	logger  *logrus.Logger
	store   *store.Store
}

func New(c *Config) *Notificator {
	return &Notificator{
		version: "0.0.1",
		conf:    c,
		logger:  logrus.New(),
	}
}

func (n *Notificator) Start() error {
	if err := n.configLogger(); err != nil {
		return err
	}
	if err := n.configStore(); err != nil {
		return err
	}
	n.logger.Info(fmt.Sprintf("<-------------------- Application version %s started -------------------->", n.version))
	return nil
}

func (n *Notificator) configLogger() error {
	level, err := logrus.ParseLevel(n.conf.LogLevel)
	if err != nil {
		return err
	}
	n.logger.SetLevel(level)
	n.logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:     false,
		FullTimestamp:   true,
		TimestampFormat: time.RFC822,
	})
	f, err := os.OpenFile(n.conf.LogName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	n.logger.SetOutput(f)
	return nil
}

func (n *Notificator) configStore() error {
	store := store.New(n.conf.Database)
	if err := store.Open(); err != nil {
		return nil
	}
	n.store = store
	return nil
}
