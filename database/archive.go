package database

import (
	"database/sql"
)

func (s *DBServer) archive() error {
	db, err := sql.Open("pgx", s.createSQLConnectionString())
	if err != nil {
		s.log.Errorf("Failed to connect to database to run archive: %v", err)
		return err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		s.log.Errorf("Failed to ping database: %v", err)
		return err
	}

	return nil
}
