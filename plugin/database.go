package plugin

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

func NewDatabase(config *Config, logger *Logger) *Database {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.GetString("database.host"),
		config.GetString("database.user"),
		config.GetString("database.password"),
		config.GetString("database.dbname"),
		config.GetString("database.port"),
		config.GetString("database.sslmode"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatal("Failed to connect to database: ", err)
	}

	logger.Info("Connected to database")

	return &Database{db}
}
