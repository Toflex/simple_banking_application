package main

import (
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"os"
	"simple_banking_application/internal/common/config"
	"simple_banking_application/internal/database"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)
}

func main() {
	config.Load()

	log.Info("Connecting to postgres")
	dbConnection := database.ConnectDB(config.Instance.DatabaseURL)
	err := database.Migration(dbConnection)
	if err != nil {
		log.Fatalf("error running db migration %s", err.Error())
	}

	defer func() {
		sqlDB, _ := dbConnection.DB()
		err = sqlDB.Close()
		if err != nil {
			log.Error("error: %v", err)
			return
		}
	}()

	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Test App v1.0.1",
	})

}
