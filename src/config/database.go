package config

import (
	"fmt"
	mdl "mahaqu/src/models"
	"mahaqu/src/utility"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnetDB() {
	dbinfo := fmt.Sprintf("sslmode=disable dbname=%s host=%s port=%s user=%s password=%s", utility.GoDotEnvVariable("DB_NAME"), utility.GoDotEnvVariable("DB_HOST"), utility.GoDotEnvVariable("DB_PORT"), utility.GoDotEnvVariable("DB_USER"), utility.GoDotEnvVariable("DB_PASSWORD"))

	db, err := gorm.Open(postgres.Open(dbinfo))

	if err != nil {
		fmt.Println(`database error`)
		panic(err)
	}

	fmt.Println(`database connected`)
	db.AutoMigrate(
		&mdl.PaymentType{},
		&mdl.DonationStatus{},
		&mdl.CategoryDonation{},
		&mdl.TargetDonation{},
		&mdl.Donation{},
	)
	DB = db
}
