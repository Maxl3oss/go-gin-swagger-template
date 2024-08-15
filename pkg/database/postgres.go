package database

import (
	"log"
	"os"
	"role-management/pkg/utils"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// PostgreSQLConnection func for connection to PostgreSQL database.
func PostgreSQLConnection() (*gorm.DB, error) {
	// Build PostgreSQL connection URL.
	connectionString, err := utils.ConnectionURLBuilder(os.Getenv("CONNECT_TYPE"))
	if err != nil {
		return nil, err
	}

	// New logger for detailed SQL logging
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Enable color
		},
	)

	// connect
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{
		Logger: newLogger, // add Logger
	})

	// if connect fail
	if err != nil {
		panic("failed to connect to database")
	}

	// AutoMigrate
	// db.Migrator().DropTable(&models.User{})
	// db.AutoMigrate(&models.User{})

	return db, nil
}
