package config

import (
	"agen_edc/internal/models"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewGormDB(cfg *Config) *gorm.DB {
	dsn := cfg.DSN()
	gcfg := &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)}
	db, err := gorm.Open(postgres.Open(dsn), gcfg)
	if err != nil {
		log.Fatalf("failed connect db: %v", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("sql db err: %v", err)
	}
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetMaxIdleConns(10)

	// AutoMigrate minimal: safe for dev. In prod gunakan migrations only.
	if err := db.AutoMigrate(
		&models.Agent{},
		&models.AcquisitionInfo{},
		&models.Owner{},
		&models.BusinessProfile{},
		&models.UploadedDocument{},
		&models.BankInfo{},
		&models.Signature{},
		&models.AuditLog{},
	); err != nil {
		log.Printf("auto migrate warning: %v", err)
	}

	return db
}
