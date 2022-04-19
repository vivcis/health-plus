package db

import (
	"fmt"

	"github.com/decadev/squad10/healthplus/models"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
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
	// SELECT * from doctor table where email = ?
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
	//MYSQL: INSERT IN patient TABLE...
	if err := DB.Create(user).Error; err != nil {
		fmt.Println(err)
		return
	}
}

func FindDocByEmailandUserName(email string, username string) (*models.Doctor, error) {
	user := &models.Doctor{}
	// SELECT * from Doctor table where email = ?
	err := DB.Where("email = ?", "username = ?", email, username).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func CreateDocInTable(user models.Doctor) {
	if err := DB.Create(user).Error; err != nil {
		fmt.Println(err)
		return
	}
}

func Authenticate(username, password string) (*models.Doctor, error) {
    //Retrieve the username and hashed password associated with the given username.
    //If matching username exists, return the ErrMismatchedHashAndPassword error.
    user := &models.Doctor{}
    err := DB.Where("username = ?", username).First(user).Error
    if err != nil {
        return nil, err
    }
    // Check whether the hashed password and plain-text password provided match
    // If they don't, we return the ErrInvalidCredentials error.
    err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
    if err != nil {
        return nil, bcrypt.ErrMismatchedHashAndPassword
    }
    return user, nil
}
func FindDoctorByUsername(username string) (*models.Doctor, error) {
    user := &models.Doctor{}
    err := DB.Where("username = ?", username).First(user).Error
    if err != nil {
        return nil, err
    }
    return user, nil
}