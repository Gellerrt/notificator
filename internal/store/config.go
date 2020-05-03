package store

type Config struct {
	// DB params
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

func NewConfig() *Config {
	return &Config{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "Aa174683Aa",
		Database: "work",
	}
}
