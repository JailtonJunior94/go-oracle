package migrate

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	oracleMigrate "github.com/golang-migrate/migrate/v4/database/oracle"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func NewMigrateOracle(db *sql.DB, migratePath, dbName string) (Migrate, error) {
	if db == nil {
		return nil, ErrDatabaseConnection
	}

	driver, err := oracleMigrate.WithInstance(db, &oracleMigrate.Config{
		MultiStmtEnabled: true,
	})
	if err != nil {
		return nil, ErrUnableToCreateDriver
	}

	migrateInstance, err := migrate.NewWithDatabaseInstance(migratePath, dbName, driver)
	if err != nil {
		return nil, err
	}
	return &migration{migrate: migrateInstance}, nil
}
