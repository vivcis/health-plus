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
	dsn := "root:flyn!GG@01@tcp(127.0.0.1:3306)/hospital?charset=utf8mb4&parseTime=True&loc=Local"
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

func FindUserByEmail(email string) (*models.Patient, error) {
	user := &models.Patient{}
	// SELECT * from patient table where email = ?
	err := DB.Where("email = ?", email).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func FindUserByUsername(username string) (*models.Patient, error) {
	user := &models.Patient{}
	// SELECT * from patient table where email = ?
	err := DB.Where("username = ?", username).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func CreatePatientInTable(user models.Patient) {
	if err := DB.Create(user).Error; err != nil {
		fmt.Println(err)
		return
	}
}

//
////Register and send patient details to DB
//func InsertPatientDetails(p models.Patient) {
//	insert, err := DataBase.Query("INSERT INTO patients (ID, NAME, AGE, EMAIL, PASSWORD) VALUES (?,?,?,?,?)", p.ID, p.Name, p.Age, p.Email, p.Password)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println(insert)
//}
//
//func ScanDoctors() []models.Doctor {
//	var Container []models.Doctor
//
//	st := "SELECT * FROM doctors"
//	rows, err := DataBase.Query(st)
//	if err != nil {
//		fmt.Println(err)
//		return nil
//	}
//
//	defer rows.Close()
//
//	for rows.Next() {
//
//		var r models.Doctor
//		err := rows.Scan(&r.ID, &r.Name, &r.Age, &r.Email, &r.PasswordHash, &r.Password, &r.Specialty, &r.WorkingHour, &r.Bookings)
//		if err != nil {
//			fmt.Println(err)
//		}
//		Container = append(Container, r)
//	}
//	return Container
//}
//
//func ScanAppoints() []models.Appointment {
//	var Container []models.Appointment
//
//	st := "SELECT * FROM appointments"
//	rows, err := DataBase.Query(st)
//	if err != nil {
//		fmt.Println(err)
//		return nil
//	}
//
//	defer rows.Close()
//
//	for rows.Next() {
//
//		var r models.Appointment
//		err := rows.Scan(&r.ID, &r.Purpose, &r.PatientID, &r.DoctorID, &r.Date, &r.AppointmentHour)
//		if err != nil {
//			fmt.Println(err)
//		}
//		Container = append(Container, r)
//	}
//	return Container
//}
//
//func ScanPatients() []models.Patient {
//	var Container []models.Patient
//
//	st := "SELECT * FROM patients"
//	rows, err := DataBase.Query(st)
//	if err != nil {
//		fmt.Println(err)
//		return nil
//	}
//
//	defer rows.Close()
//
//	for rows.Next() {
//
//		var r models.Patient
//		err := rows.Scan(&r.ID, &r.Name, &r.Age, &r.Email, &r.PasswordHash, &r.Password, &r.Appointments)
//		if err != nil {
//			fmt.Println(err)
//		}
//		Container = append(Container, r)
//	}
//	return Container
//}
