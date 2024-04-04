package main

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"simple_banking_application/internal/models"
)

func seeder(tx *gorm.DB) {
	users, err := seedUser(tx)
	if err != nil {
		log.Fatalf("error seeding users. %s", err.Error())
	}

}

func seedUser(tx *gorm.DB) ([]models.User, error) {
	users := []models.User{
		{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "JohnDoe@test.com",
			Account: models.Account{
				Currency: models.NGN,
				Balance:  0,
			},
		},
		{
			FirstName: "Peter",
			LastName:  "Rock",
			Email:     "PeterRock@mail.com",
			Account: models.Account{
				Currency: models.NGN,
				Balance:  0,
			},
		},
	}

	if err := tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
