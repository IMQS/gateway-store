package database

import (
	"errors"
	"strconv"
	"strings"

	"github.com/IMQS/gateway-store/config"
	"github.com/IMQS/log"
	"github.com/jackc/pgx"
)

// DBServer contains the configuration, db connection pool and logging for the package
//
// Why 'ready' channel?
// We run DB migrations in a goroutine because we can't block startup of the service (it times out).
// We need a way to know when DB migrations are done, so we send notification via the 'ready' channel.
//
type DBServer struct {
	config   *config.DB
	connPool *pgx.ConnPool
	log      *log.Logger
	Ready    chan bool
}

// NewDBServer returns a DB structs and also setups the DB and runs migrations
func NewDBServer(conf *config.DB, log *log.Logger) (*DBServer, error) {
	s := &DBServer{
		config:   conf,
		connPool: nil,
		log:      log,
		Ready:    make(chan bool),
	}

	err := s.createPGXConnectionPool()
	if err != nil {
		return nil, err
	}

	//Run the migrate in the background
	s.migrate(log)

	log.Debug("DB Server Created")
	return s, nil
}

func (s *DBServer) runArchive() {
	s.log.Infof("Running Database Archive")
	if err := s.archive(); err != nil {
		s.log.Errorf("Database migration: Migration failed (%s)", err)
		s.Ready <- false
		return
	}
	s.log.Infof("Database migration: Done, service ready to accept requests")
	s.Ready <- true
}

// Close disconnects from the database
func (s *DBServer) Close() error {
	s.connPool.Close()
	return nil
}

// buildInsertQueryBuilder creates an INSERT statement into `table` using
// `fields` to specify the specific fields to be inserted
func buildInsertQueryBuilder(table string, fields []string, returnRowID bool) string {
	var builder strings.Builder
	builder.Reset()
	builder.WriteString("INSERT INTO " + table + " (")
	i := 0
	for _, field := range fields {
		if i != 0 {
			builder.WriteString(", ")
		}
		if field == "geometry" {
			// TODO: when we start handling it in the API
		} else {
			builder.WriteString(field)
		}
		i++
	}
	builder.WriteString(") VALUES (")
	for j := range fields {
		if j != 0 {
			builder.WriteString(", ")
		}
		builder.WriteString("$")
		builder.WriteString(strconv.Itoa(j + 1))
	}
	builder.WriteString(")")
	if returnRowID {
		builder.WriteString(" RETURNING _id")
	}
	return builder.String()
}

func buildUpdateQueryBuilder(table string, fields []string, whereField string) string {
	var builder strings.Builder
	builder.Reset()
	builder.WriteString("UPDATE " + table + " SET ")
	i := 1
	for _, field := range fields {
		if i != 1 {
			builder.WriteString(", ")
		}
		if field == "geometry" {
			// TODO: when we start handling it in the API
		} else {
			builder.WriteString(field + " = $" + strconv.Itoa(i))
		}
		i++
	}
	if whereField != "" {
		builder.WriteString(" WHERE " + whereField + " = $" + strconv.Itoa(i))
	}
	return builder.String()
}

// ConflictAction is an enum used for BuildOnConflict
type ConflictAction int

// This is the closest thing to enums in go :(
const (
	CAUnknown ConflictAction = iota
	CANothing
	CAUpdate
)

// buildOnConflict creates an ON CONFLICT statement using `conflicts` as the
// conflict fields (an empty array of `conflicts` will be appropriately
// handled). `action` will determine the `conflict_action` that is created.
// `fields` will be used for a `conflict_action` that requires fields.
func buildOnConflict(conflicts []string, action ConflictAction, fields []string) (string, error) {
	var builder strings.Builder
	builder.Reset()
	builder.Grow(2048)
	builder.WriteString(" ON CONFLICT")

	if len(conflicts) > 0 {
		builder.WriteString(" (")
		for i, field := range conflicts {
			if i != 0 {
				builder.WriteString(", ")
			}
			builder.WriteString(field)
		}
		builder.WriteString(")")
	}

	builder.WriteString(" DO ")
	switch action {
	case CANothing:
		builder.WriteString("NOTHING")
	case CAUpdate:
		builder.WriteString(buildUpdateQueryBuilder("", fields, ""))
	default:
		return "", errors.New("unknown `conflict_action` for `ON CONFLICT`")
	}

	return builder.String(), nil
}
