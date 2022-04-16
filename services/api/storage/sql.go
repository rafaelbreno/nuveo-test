package storage

import (
	"fmt"

	"github.com/rafaelbreno/nuveo-test/internal"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type (
	// SQL stores connection
	// and config for dealing
	// with SQL databases.
	SQL struct {
		Client *gorm.DB
		in     *internal.Internal
	}
)

var (
	pgDSN = `host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`
)

// NewSQL returns an instance of SQL based
// on given Internal struct.
func NewSQL(in *internal.Internal) (*SQL, error) {
	dsn := fmt.Sprintf(
		pgDSN,
		in.Cfg.Database.PGHost,
		in.Cfg.Database.PGPort,
		in.Cfg.Database.PGUser,
		in.Cfg.Database.PGPassword,
		in.Cfg.Database.PGDBName,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		in.L.Error(err.Error())
		return &SQL{}, err
	}

	return &SQL{
		in:     in,
		Client: db,
	}, err
}
