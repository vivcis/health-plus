package db

import (
	"fmt"
	"github.com/decadev/squad10/healthplus/models"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func SetupDB() {
	password := os.Getenv("DB_PASSWORD")
	dbDatabase := os.Getenv("DB")
	root := os.Getenv("DB_ROOTS")
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", root, password, dbDatabase)
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
	// SELECT * from Patient table where email = ?
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

func FindDoctorByID(id string) *models.Doctor {
	user := &models.Doctor{}
	err := DB.Where("id = ?", id).First(user).Error
	if err != nil {
		return nil
	}
	return user
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

func AuthenticatePatient(username, password string) (*models.Patient, error) {
	//Retrieve the username and hashed password associated with the given username.
	//If matching username exists, return the ErrMismatchedHashAndPassword error.
	user := &models.Patient{}
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
func FindPatientByUsername(username string) (*models.Patient, error) {
	user := &models.Patient{}

	err := DB.Where("username = ?", username).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetAllDoctors() []models.Doctor {
	var users []models.Doctor
	sqlDB, _ := DB.DB()

	st := "SELECT * FROM doctors"
	rows, err := sqlDB.Query(st)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	for rows.Next() {
		var r models.Doctor
		err := rows.Scan(&r.ID, &r.Username, &r.Name, &r.Age, &r.Email, &r.PasswordHash, &r.Specialty, &r.WorkingHour)
		if err != nil {
			fmt.Println(err)
		}
		users = append(users, r)
	}

	return users
}

func FindDoctorByIDandUserName(id string) *models.Doctor {
	user := &models.Doctor{}
	// SELECT * from doctor table where email = ?
	err := DB.Where("id = ?", id).First(user).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return user
}

func CreateAppointmentInTable(user models.Appointment) {
	//MYSQL: INSERT IN patient TABLE...
	if err := DB.Create(&user).Error; err != nil {
		fmt.Println(err)
		return
	}
}

func FindPatientAppointment(id string) []models.Appointment {
	sqlDB, _ := DB.DB()

	rows, err := sqlDB.Query("SELECT * FROM appointments WHERE patient_id = ?", id)
	if err != nil {
		fmt.Println(err)
	}

	var users []models.Appointment

	for rows.Next() {
		var user models.Appointment
		err := rows.Scan(&user.ID, &user.Purpose, &user.PatientID, &user.DoctorID, &user.Date, &user.AppointmentHour, &user.DoctorName, &user.PatientName)
		if err != nil {
			fmt.Println(err)
		}
		users = append(users, user)
	}
	return users
}

func DeleteAppointmentbyID(id string) {
	sqlDB, _ := DB.DB()

	del, err := sqlDB.Prepare("DELETE FROM appointments WHERE (id = ?);")
	if err != nil {
		fmt.Println(err)
	}
	defer del.Close()
	del.Exec(id)
}

func FindDoctorAppointment(id string) []models.Appointment {
	sqlDB, _ := DB.DB()

	rows, err := sqlDB.Query("SELECT * FROM appointments WHERE doctor_id = ?", id)
	if err != nil {
		fmt.Println(err)
	}

	var users []models.Appointment

	for rows.Next() {
		var user models.Appointment
		err := rows.Scan(&user.ID, &user.Purpose, &user.PatientID, &user.DoctorID, &user.Date, &user.AppointmentHour, &user.DoctorName, &user.PatientName)
		if err != nil {
			fmt.Println(err)
		}
		users = append(users, user)
	}
	return users
}
