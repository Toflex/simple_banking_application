package database

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"simple_banking_application/internal/models"
)

var gormDB *gorm.DB

func ConnectDB(dsn string) *gorm.DB {
	dialector := postgres.New(postgres.Config{
		DSN: dsn,
	})
	db, err := gorm.Open(dialector, &gorm.Config{
		PrepareStmt: true,
	})

	if err != nil {
		log.Fatalf("failed to connect database. %s", err.Error())
	}

	setDB(db)

	return db
}

func Migration(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
		&models.Account{},
		&models.Transaction{})
}

func setDB(db *gorm.DB) {
	gormDB = db
}

func DB() *gorm.DB {
	return gormDB
}
