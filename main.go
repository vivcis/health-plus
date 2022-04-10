package main

import (
	"fmt"
	"github.com/decadev/squad10/healthplus/DBConnections"
	"github.com/decadev/squad10/healthplus/handlers"
	"github.com/decadev/squad10/healthplus/models"
	"github.com/go-chi/chi"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

func main() {
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

	DBConnections.OpenDB()

	router := chi.NewRouter()
	fmt.Println("Server up and Running")

	handlers.Register(router)

	e := http.ListenAndServe(":8080", router)
	if e != nil {
		fmt.Println(e)
		return
	}

}
