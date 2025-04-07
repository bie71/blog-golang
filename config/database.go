package config

import (
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Postgres struct {
	DB *gorm.DB
}

func (cfg Config) ConectionPostgres() (*Postgres, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta",
		cfg.DB.DBHost,
		cfg.DB.DBPort,
		cfg.DB.DBUser,
		cfg.DB.DBPassword,
		cfg.DB.DBName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to connect to database")
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Error().Err(err).Msg("Failed to get database connection")
		return nil, err
	}
	sqlDB.SetMaxIdleConns(cfg.DB.DBMaxIdle)
	sqlDB.SetMaxOpenConns(cfg.DB.DBMaxOpen)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.DB.DBMaxLife) * time.Second)

	return &Postgres{DB: db}, nil
}
