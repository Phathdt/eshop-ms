package component

import (
	"fmt"
	"os"
	"time"

	"user_api/pkg/config"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	gormLogger "gorm.io/gorm/logger"
)

func NewGormService() (*gorm.DB, error) {
	cfg := config.Config

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", cfg.POSTGRES.Host, cfg.POSTGRES.User, cfg.POSTGRES.Pass, cfg.POSTGRES.Database, cfg.POSTGRES.Port, cfg.POSTGRES.Sslmode)

	gLogger := initLogger(cfg.App.LogLevel, cfg.GORM.JsonLogger)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{
		Logger: gLogger,
	})

	if err != nil {
		return nil, err
	}

	sqlConn, _ := db.DB()

	if err = sqlConn.Ping(); err != nil {
		return nil, fmt.Errorf("db.Ping %w", err)
	}

	return db, nil
}

func initLogger(level string, jsonLogger bool) gormLogger.Interface {
	var logLevel gormLogger.LogLevel

	switch level {
	case "INFO":
		logLevel = gormLogger.Info
	case "FATAL":
		logLevel = gormLogger.Error
	case "ERROR":
		logLevel = gormLogger.Error
	case "WARN":
		logLevel = gormLogger.Warn
	case "DEBUG":
		logLevel = gormLogger.Info
	case "PANIC":
		logLevel = gormLogger.Error
	default:
		logLevel = gormLogger.Info
	}

	log := logrus.New()
	colorful := true

	if jsonLogger {
		log.SetFormatter(&logrus.JSONFormatter{})
		colorful = false
	}

	log.SetOutput(os.Stdout)

	newLogger := gormLogger.New(log, gormLogger.Config{
		SlowThreshold:             time.Second, // Slow SQL threshold
		IgnoreRecordNotFoundError: false,       // Skip ErrRecordNotFound error for logger
		LogLevel:                  logLevel,    // Log level. Default value: gormLogger.Info
		Colorful:                  colorful,
	})

	return newLogger
}
