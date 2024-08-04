package database

import (
	"context"

	"github.com/travel-api-main/domain/models"
	"github.com/travel-api-main/pkg/config"
	"github.com/travel-api-main/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func PostgresConnection(ctx context.Context) {
	l := logger.GetLoggerContext(ctx, "database", "Connect")
	dsn := config.GetString("postgres_dsn")

	// l.Info(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	l.Info("open connection to postgres")

	// db.AutoMigrate(&models.User{})
	// db.AutoMigrate(&models.Token{})
	db.AutoMigrate(&models.Country{})
	db.AutoMigrate(&models.Travel{})

	DB = db
}
