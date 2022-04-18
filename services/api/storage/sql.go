package storage

import (
	"fmt"
	"time"

	"github.com/rafaelbreno/nuveo-test/entity"
	"github.com/rafaelbreno/nuveo-test/internal"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	var db *gorm.DB
	var err error

	for i := int64(1); i <= 5; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
		if err != nil {
			in.L.Error(err.Error())
			time.Sleep(time.Second * time.Duration(i*2))
			continue
		}
		break
	}

	if err != nil {
		return &SQL{}, err
	}

	in.L.Info("connected to db")

	s := &SQL{
		in:     in,
		Client: db,
	}

	if err := s.runMigrations(); err != nil {
		in.L.Error(err.Error())
		return &SQL{}, err
	}

	return s, nil
}

func (s *SQL) runMigrations() error {
	return s.
		Client.
		AutoMigrate(entity.User{})
}
