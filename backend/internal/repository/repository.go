package repository

import (
	"errors"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/seu-usuario/minha-api/internal/config"
)

/*--------------------------------------------------------*/
func InitRepo(cfg config.Config) (*gorm.DB, error) {
	if err := runMigrations(cfg); err != nil {
		return nil, err
	}

	return gorm.Open(postgres.Open(cfg.GetDbDsn()), &gorm.Config{})
}

/*--------------------------------------------------------*/
func runMigrations(cfg config.Config) error {
	m, err := migrate.New("file://migrations", cfg.GetDbUrl())
	if err != nil {
		return err
	}
	defer m.Close()

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}
