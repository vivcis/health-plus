package db

import (
	"fmt"
	"github.com/decadev/squad10/healthplus/models"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDB() {
	dsn := "root:appliCATION123@#@tcp(127.0.0.1:3306)/hospital?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}

	//This creates our table for this model
	err = db.AutoMigrate(&models.Doctor{}, &models.Patient{}, &models.Appointment{})
	if err != nil {
		fmt.Println(err)
		return
	}
	DB = db
}

func FindUserByEmailandUserName(email string, username string) (*models.Patient, error) {
	user := &models.Patient{}
	// SELECT * from patient table where email = ?
	err := DB.Where("email = ?", email).First(user).Error
	if err != nil {
		err = DB.Where("username = ?", username).First(user).Error
		if err != nil {
			return nil, err
		}
	}
	return user, nil
}

func CreatePatientInTable(user models.Patient) {
	if err := DB.Create(user).Error; err != nil {
		fmt.Println(err)
		return
	}
}
