package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx"
)

func (s *DBServer) createPGXConnectionPool() error {
	var err error
	connStr := s.createSQLConnectionString()
	s.connPool, err = s.connectPool(connStr)
	if err != nil {
		if err.(pgx.PgError).Code == "3D000" {
			err = s.createDatabase(connStr)
			if err != nil {
				return err
			}
			s.connPool, err = s.connectPool(connStr)
			if err != nil {
				s.log.Errorf("Failed to migrate connection pool: %v", err)
				return err
			}
		} else {
			s.log.Errorf("Failed to migrate connection pool: %v", err)
			return err
		}
	}
	return err
}

func (s *DBServer) connectSingle(config *pgx.ConnConfig) (*pgx.Conn, error) {
	conn, err := pgx.Connect(*config)
	if err != nil {
		s.log.Errorf("Failed to connect to database: %v", err)
		return nil, err
	}

	err = conn.Ping(context.Background())
	if err != nil {
		s.log.Errorf("Failed to ping database: %v", err)
		conn.Close()
		return nil, err
	}
	return conn, nil
}

func (s *DBServer) connectPool(connStr string) (*pgx.ConnPool, error) {
	pgxConfig, err := pgx.ParseDSN(connStr)
	if err != nil {
		s.log.Errorf("Failed to parse PGX DSN string: %v", err)
		return nil, err
	}

	conn, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig:     pgxConfig,
		MaxConnections: 5,
	})
	return conn, err
}

func (s *DBServer) createDatabase(connStr string) error {
	pgxConfig, err := pgx.ParseDSN(connStr)
	pgxConfig.Database = "postgres"

	conn, err := s.connectSingle(&pgxConfig)
	if err != nil {
		return err
	}
	_, err = conn.Exec(fmt.Sprintf("CREATE DATABASE %s OWNER imqs;", s.config.Dbname))
	if err != nil {
		s.log.Errorf("Failed to migrate database '%s': %v", s.config.Dbname, err)
		conn.Close()
		return err
	}
	conn.Close()

	pgxConfig.Database = s.config.Dbname
	conn, err = s.connectSingle(&pgxConfig)
	if err != nil {
		return err
	}

	_, err = conn.Exec("CREATE EXTENSION IF NOT EXISTS postgis;")
	if err != nil {
		s.log.Errorf("Could not migrate postgis extension: %v", err)
	}
	conn.Close()
	return err
}

func (s *DBServer) createSQLConnectionString() string {
	return fmt.Sprintf("user=%s password=%s database=%s host=%s port=%s sslmode=disable", s.config.Username,
		s.config.Password, s.config.Dbname, s.config.Host, s.config.Port)
}
