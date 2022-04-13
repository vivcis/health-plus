package handlers

import (
	"fmt"
	"github.com/decadev/squad10/healthplus/db"
	"github.com/decadev/squad10/healthplus/models"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
	"strconv"
)

//Indexhandler gets the homepage
func Indexhandler(w http.ResponseWriter, r *http.Request) {
	t, e := template.ParseFiles("pages/index.html")
	if e != nil {
		fmt.Println(e)
		return
	}
	e = t.Execute(w, nil)
	if e != nil {
		fmt.Println(e)
		return
	}
}

// RegisterPatientHandler gets Patient's SignUp page
func RegisterPatientHandler(w http.ResponseWriter, r *http.Request) {
	t, e := template.ParseFiles("pages/registerPatient.html")
	if e != nil {
		fmt.Println(e)
		return
	}
	e = t.Execute(w, nil)
	if e != nil {
		fmt.Println(e)
		return
	}
}

//PostRegisterPatientHandler successfully register's patient's name in the db if valid
func PostRegisterPatientHandler(w http.ResponseWriter, r *http.Request) {
	var user models.Patient
	r.ParseForm()
	name := r.FormValue("name")
	ageString := r.FormValue("ageString")
	email := r.FormValue("email")
	username := r.FormValue("username")
	password := r.FormValue("password")

	age, _ := strconv.Atoi(ageString)

	user.ID = uuid.NewString()
	user.Name = name
	user.Age = uint(age)
	user.Email = email
	user.Username = username
	user.Password = password

	_, err := db.FindUserByEmailandUserName(user.Email, user.Username)
	if err == nil {
		// this user already exists
		// return a message to the user

		t, e := template.ParseFiles("pages/registerPatient.html")
		if e != nil {
			fmt.Println(e)
			return
		}
		e = t.Execute(w, "User already exists. Check Email or Username")
		if e != nil {
			fmt.Println(e)
			return
		}

	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		return
	}
	user.PasswordHash = string(hashedPassword)

	db.CreatePatientInTable(user)

	http.Redirect(w, r, "/patientLogin", http.StatusFound)

}

func PatientLoginHandler(w http.ResponseWriter, r *http.Request) {
	t, e := template.ParseFiles("pages/patientLogin.html")
	if e != nil {
		fmt.Println(e)
		return
	}
	e = t.Execute(w, nil)
	if e != nil {
		fmt.Println(e)
		return
	}
}
