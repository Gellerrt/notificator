package server

type Config struct {
	// our server params
	Host     string `yaml:"server_host"`
	Port     string `yaml:"server_port"`
	URI      string `yaml:"uri_notifications"`
	LogName  string `yaml:"server_log_name"`
	LogLevel string `yaml:"server_log_level"`
}

func NewConfig() *Config {
	return &Config{
		Host:     "localhost",
		Port:     "8000",
		URI:      "/notificate",
		LogName:  "server.log",
		LogLevel: "DEBUG",
	}
}
