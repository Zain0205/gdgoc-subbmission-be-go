package database

import (
	"log"

	"github.com/Zain0205/gdgoc-subbmission-be-go/config"
	"github.com/Zain0205/gdgoc-subbmission-be-go/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	dsn := config.DB_DSN

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!", err)
	}

	log.Println("Database connection established.")
}

func MigrateDatabase() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Track{},
		&models.Series{},
		&models.Submission{},
		&models.Leaderboard{},
		&models.UserSeriesVerification{},
		&models.AchievementType{},
		&models.Achievement{},
		&models.UserAchievement{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database!", err)
	}
	log.Println("Database migrated.")
}

