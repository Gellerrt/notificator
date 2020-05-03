package config

type Config struct {

	// loggers
	LogName       string `yaml:"log_name"`
	LogLevel      string `yaml:"log_level"`
	ServerLogName string `yaml:"server_log_name"`
	ServerLogLevel string `yaml:"server_log_level"`

	// where to send
	URL      string `yaml:"URL"`
	PortSend string `yaml:"port_send"`

	// DB params
	HostDB     string `yaml:"host"`
	PortDB     string `yaml:"port"`
	UserDB     string `yaml:"user"`
	PasswordDB string `yaml:"password"`

	// our server params
	ServerHost string `yaml:"server_host"`
	ServerPort string `yaml:"server_port"`
	URINotificate string `yaml:"uri_notifications"`

	// crons
	CronSend       string `yaml:"cron"`
	CronAddMessage string `yaml:"cron_add_message"`
}

func NewConfig() *Config {
	return &Config{
		LogName:        "logger",
		LogLevel:       "DEBUG",
		ServerLogName:  "server",
		ServerLogLevel: "DEBUG",
		URL:            "http:\\localhost",
		PortSend:       "8000",
		HostDB:         "localhost",
		ServerHost:     "localhost",
		ServerPort:     "8000",
		URINotificate: "/notificate",
	}
}
