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
	var p models.Patient
	r.ParseForm()
	name := r.FormValue("name")
	ageString := r.FormValue("ageString")
	email := r.FormValue("email")
	username := r.FormValue("username")
	password := r.FormValue("password")

	age, _ := strconv.Atoi(ageString)

	p.ID = uuid.NewString()
	p.Name = name
	p.Age = uint(age)
	p.Email = email
	p.Username = username
	p.Password = password

	_, err := db.FindUserByEmail(p.Email)
	if err == nil {
		// this user already exists
		// return a message to the user

		t, e := template.ParseFiles("pages/registerPatient.html")
		if e != nil {
			fmt.Println(e)
			return
		}
		e = t.Execute(w, "email already exist")
		if e != nil {
			fmt.Println(e)
			return
		}

	}

	_, err = db.FindUserByUsername(p.Username)
	if err == nil {
		// this username already exists
		// return a message to the user

		t, e := template.ParseFiles("pages/registerPatient.html")
		if e != nil {
			fmt.Println(e)
			return
		}
		e = t.Execute(w, "username already in use")
		if e != nil {
			fmt.Println(e)
			return
		}

	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		return
	}
	p.PasswordHash = string(hashedPassword)

	db.CreatePatientInTable(p)

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
