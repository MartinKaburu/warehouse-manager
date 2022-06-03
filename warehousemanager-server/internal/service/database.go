package service

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Name     string
	Password string
}

func ConnectToPostgres(dbConfig string) (*gorm.DB, error) {
	gormConfig := &gorm.Config{}
	gormConfig.Logger = logger.Default.LogMode(logger.Silent)
	db, err := gorm.Open(postgres.Open(dbConfig), gormConfig)
	if err != nil {
		return nil, err
	}

	return db, nil
}
