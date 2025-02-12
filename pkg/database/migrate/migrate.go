package migrate

import (
	"errors"

	"github.com/golang-migrate/migrate/v4"
)

var (
	ErrMigrateVersion       = errors.New("error on migrate version")
	ErrDatabaseConnection   = errors.New("database connection is nil")
	ErrUnableToCreateDriver = errors.New("unable to create driver instance")
)

type (
	Migrate interface {
		Execute() error
	}

	migration struct {
		migrate *migrate.Migrate
	}
)

func (m *migration) Execute() error {
	err := m.migrate.Up()
	if err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			return nil
		}
		return err
	}

	return nil
}
