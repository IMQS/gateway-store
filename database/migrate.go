package database

import (
	"database/sql"
	"github.com/IMQS/log"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	"github.com/golang-migrate/migrate/source/go_bindata"
	_ "github.com/jackc/pgx/stdlib"
)

func (s *DBServer) migrate(log *log.Logger) error {

	log.Debug("Running migrate")
	db, err := sql.Open("pgx", s.createSQLConnectionString())
	if err != nil {
		s.log.Errorf("Failed to connect to database to run migrate scripts: %v", err)
		return err
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		s.log.Errorf("Failed to ping database: %v", err)
		return err
	}

	source := bindata.Resource(AssetNames(),
		func(name string) (bytes []byte, e error) {
			return Asset(name)
		})

	// Setup a bindata source
	data, err := bindata.WithInstance(source)
	if err != nil {
		s.log.Errorf("Failed to set migrate bindata source: %v", err)
		return err
	}

	// Setup a migrate driver to postgres
	dbDriver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		s.log.Errorf("failed to set migrate database driver: %v", err)
		return err
	}

	// Create migrate instance
	migration, err := migrate.NewWithInstance("go-bindata", data, "postgres", dbDriver)
	if err != nil {
		s.log.Errorf("Failed to migrate migration instance: %v", err)
		return err
	}

	err = migration.Up()
	if err != nil {
		if err == migrate.ErrNoChange {
			s.log.Infof("Database migration: No new migrations")
			err = nil
		} else {
			s.log.Errorf("Failed to run migrations: %v", err)
		}
		return err
	}

	// Run archive in goroutine to not block startup of service
	// API requests will, however, block until migrations are done
	go s.runArchive()

	log.Debug("Migrate Done")
	return nil
}
