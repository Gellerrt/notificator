package notificator

import "notificator/internal/store"

type Config struct {
	// loggers
	LogName  string `yaml:"log_name"`
	LogLevel string `yaml:"log_level"`

	// where to send
	URL      string `yaml:"URL"`
	PortSend string `yaml:"port_send"`

	// crons
	CronSend       string `yaml:"cron"`
	CronAddMessage string `yaml:"cron_add_message"`

	// store
	Database *store.Config
}

func NewConfig() *Config {
	return &Config{
		LogName:  "logger",
		LogLevel: "DEBUG",
		URL:      "http:\\localhost",
		PortSend: "8000",
		Database: store.NewConfig(),
	}
}
