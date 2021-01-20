package database

import (
	"github.com/pratheeshpcplpta/simple-meeting-scheduler/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB
var err error

func InitConnection() *gorm.DB {
	db, err = gorm.Open(sqlite.Open("simple-meeting-scheduler.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

//
// Migrate models
//
func MigrateModels() {
	db := InitConnection()
	// Migrate the schema
	db.AutoMigrate(&models.Users{})
	db.AutoMigrate(&models.MeetingSchedules{})
	db.AutoMigrate(&models.UserMeetings{})
}

//
// GetUserByUsername returns the user that the given username corresponds to. If no user is found, an
// error is thrown.
func GetUserByUsername(username string) (models.Users, error) {
	u := models.Users{}
	err := db.Where("username = ?", username).First(&u).Error
	return u, err
}
