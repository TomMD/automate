package postgres

import (
	gosql "database/sql"
	"time"

	"github.com/go-gorp/gorp"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/chef/automate/components/applications-service/pkg/config"
	"github.com/chef/automate/lib/db/migrator"
	"github.com/chef/automate/lib/logger"
)

// Postgres is a wrapping struct that will hold the database mapping object
// from the underlying db/sql implementation (gorp) plus our service config
// specifically for storage.
//
// Additionally this struct implements our storage.Client interface
type Postgres struct {
	*gorp.DbMap
	*config.Postgres
}

// New creates a new Postgres client
func New(dbConf *config.Postgres) (*Postgres, error) {
	pg := &Postgres{Postgres: dbConf}

	log.WithFields(log.Fields{
		"uri": pg.URI,
	}).Debug("Connecting to PostgreSQL")
	err := pg.connect()
	if err != nil {
		return pg, err
	}

	log.WithFields(log.Fields{
		"uri":    pg.URI,
		"schema": pg.SchemaPath,
	}).Debug("Initializing database")
	err = pg.initDB()
	return pg, err
}

// ping will verify if the database mapped with gorp is available
func (db *Postgres) ping() error {
	return db.Db.Ping()
}

// connect opens a connection to the database
func (db *Postgres) connect() error {
	dbMap, err := gosql.Open("postgres", db.URI)
	if err != nil {
		return errors.Wrapf(err, "Failed to open database with uri: %s", db.URI)
	}

	// Configure the database mapping object
	db.DbMap = &gorp.DbMap{Db: dbMap, Dialect: gorp.PostgresDialect{}}

	// Verify database
	err = db.ping()
	if err != nil {
		return errors.Wrapf(err, "Failed to ping database with uri: %s", db.URI)
	}

	return nil
}

// initDB initializes the database
func (db *Postgres) initDB() error {
	// Create the schema
	// @afiune Can we rename this library?
	// @sr Just do it ;)
	if err := migrator.Migrate(db.URI, db.SchemaPath, logger.NewLogrusStandardLogger(), false); err != nil {
		return errors.Wrapf(err, "Unable to create database schema. [path:%s]", db.SchemaPath)
	}

	// Add the tables to the database mappings
	db.AddTableWithName(deployment{}, "deployment").SetKeys(true, "id")
	db.AddTableWithName(supervisor{}, "supervisor").SetKeys(true, "id")
	db.AddTableWithName(serviceGroup{}, "service_group").SetKeys(true, "id")
	db.AddTableWithName(service{}, "service").SetKeys(true, "id")

	//return db.CreateTablesIfNotExists() // I don't think we can ensure the foreign keys
	return nil
}

type dbTable interface {
	NeedUpdate() bool
}

// service struct is the representation of the service table inside the db
type service struct {
	ID                  int32     `db:"id"`
	Origin              string    `db:"origin"`
	Name                string    `db:"name"`
	Version             string    `db:"version"`
	Release             string    `db:"release"`
	Status              string    `db:"status"`
	Health              string    `db:"health"`
	GroupID             int32     `db:"group_id"`
	DeploymentID        int32     `db:"deployment_id"`
	SupID               int32     `db:"sup_id"`
	Channel             string    `db:"channel"`
	FullPkgIdent        string    `db:"package_ident"`
	PreviousHealth      string    `db:"previous_health"`
	UpdateStrategy      string    `db:"update_strategy"`
	LastEventOccurredAt time.Time `db:"last_event_occurred_at"`
	HealthUpdatedAt     time.Time `db:"health_updated_at"`
	CreatedAt           time.Time `db:"-"`
	UpdatedAt           time.Time `db:"-"`

	// (internal) use it to know if the service needs an update or not
	needUpdate bool `db:"-"`
}

func (svc *service) NeedUpdate() bool {
	return svc.needUpdate
}

// supervisor struct is the representation of the supervisor table inside the db
type supervisor struct {
	ID        int32     `db:"id"`
	MemberID  string    `db:"member_id"`
	Fqdn      string    `db:"fqdn"`
	Site      string    `db:"site"`
	CreatedAt time.Time `db:"-"`
	UpdatedAt time.Time `db:"-"`

	// (internal) use it to know if the supervisor needs an update or not
	needUpdate bool `db:"-"`
}

func (s *supervisor) NeedUpdate() bool {
	return s.needUpdate
}

// serviceGroup struct is the representation of the service_group table inside the db
type serviceGroup struct {
	ID           int32     `db:"id"`
	Name         string    `db:"name"`
	DeploymentID int32     `db:"deployment_id"`
	CreatedAt    time.Time `db:"-"`
	UpdatedAt    time.Time `db:"-"`

	// (internal) use it to know if the service-group needs an update or not
	needUpdate bool `db:"-"`
}

func (sg *serviceGroup) NeedUpdate() bool {
	return sg.needUpdate
}

// deployment struct is the representation of the deployment table inside the db
type deployment struct {
	ID          int32     `db:"id"`
	AppName     string    `db:"app_name"`
	Environment string    `db:"environment"`
	CreatedAt   time.Time `db:"-"`
	UpdatedAt   time.Time `db:"-"`

	// (internal) use it to know if the deployment needs an update or not
	needUpdate bool `db:"-"`
}

func (d *deployment) NeedUpdate() bool {
	return d.needUpdate
}
