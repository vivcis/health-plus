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


//-------------------------RegisterDoctorHandler gets Doctor's SignUp page-----------------------------------------------
func RegisterDoctorHandler(w http.ResponseWriter, r *http.Request) {
	t, e := template.ParseFiles("pages/doctorRegister.html")
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

//-------------------PostRegisterDoctorHandler successfully registers doctor's name in the db if valid----------------------------
func PostRegisterDoctorHandler(w http.ResponseWriter, r *http.Request) {
	var user models.Doctor

	r.ParseForm()
	name := r.FormValue("name")
	ageString := r.FormValue("ageString")
	email := r.FormValue("email")
	username := r.FormValue("username")
	password := r.FormValue("password")
	speciality := r.FormValue("speciality")

	age, _ := strconv.Atoi(ageString)

	user.ID = uuid.NewString()
	user.Name = name
	user.Age = uint(age)
	user.Email = email
	user.Username = username
	user.Password = password
	user.Specialty = speciality

	_, err := db.FindDocByEmailandUserName(user.Email, user.Username)
	if err == nil {


		t, e := template.ParseFiles("pages/doctorRegister.html")
		if e != nil {
			fmt.Println(e)
			return
		}
		e = t.Execute(w, "User already exists, confirm email or username")
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

	db.CreateDocInTable(user)

	http.Redirect(w, r, "/doctorLogin", http.StatusFound)

}

//------------------------------DoctorLoginHandler gets Doctor's Login page---------------------------------------------------------
func DoctorLoginHandler(w http.ResponseWriter, r *http.Request) {
	t, e := template.ParseFiles("pages/doctorLogin.html")
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