package store

import (
	"context"
	"fmt"
	"github.com/jackc/pgx"
)

type Store struct {
	config *Config
	conn   *pgx.Conn
}

func New(conf *Config) *Store {
	return &Store{
		config: conf,
	}
}

func (s *Store) Open() error {
	connection, err := s.connectToDatabaseFromConfig()
	if err != nil {
		return err
	}
	if err := connection.Ping(context.Background()); err != nil { // TODO context
		return err
	}
	s.conn = connection
	return nil
}

func (s *Store) Close() {
	s.conn.Close()
}

// create connection to store
func (s *Store) connectToDatabaseFromConfig() (*pgx.Conn, error) {
	conf, err := pgx.ParseConnectionString(fmt.Sprintf(
		"port=%s host=%s user=%s password=%s dbname=%s sslmode=disable",
		s.config.Port, s.config.Host, s.config.User, s.config.Password, s.config.Database))
	if err != nil {
		return nil, err
	}
	conn, err := pgx.Connect(conf)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
