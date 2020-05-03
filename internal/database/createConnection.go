package database

import (
	"context"
	"fmt"
	l4g "github.com/alecthomas/log4go"
	"github.com/jackc/pgx"
	"github.com/magiconair/properties"
	"notificator/notificator/internal/config"
	"os"
)

// initialize connection to database
func InitDb(props *properties.Properties, log *l4g.Logger) *pgx.Conn {
	host, port, database, user, password := initDbParams(props, log)
	conn := connectToDatabaseFromConfig(host, port, database, user, password, log)
	log.Info("Successfully initialized connection to database")
	return conn
}

// initialize parameters to connect to database
func initDbParams(props *properties.Properties, log *l4g.Logger) (string, string, string, string, string) {
	host := config.ParseField(config.HOST, props)
	port := config.ParseField(config.PORT, props)
	database := config.ParseField(config.DATABASE, props)
	user := config.ParseField(config.USER, props)
	password := config.ParseField(config.PASSWORD, props)
	log.Info(fmt.Sprintf("Got parameters to init connection to database:\n"+
		"host=%s; port=%s; database=%s; user=%s; password=%s",
		host, port, database, user, password))
	return host, port, database, user, password
}

// create connection to database
func connectToDatabaseFromConfig(host, port, database, user, password string, log *l4g.Logger) *pgx.Conn {
	config, err := pgx.ParseConfig(fmt.Sprintf(
		"port=%s host=%s user=%s password=%s dbname=%s sslmode=disable",
		port, host, user, password, database))
	conn, err := pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		_ = log.Error(fmt.Sprintf("Unable to connect to database: %v\n", err))
		os.Exit(1)
	}
	return conn
}
