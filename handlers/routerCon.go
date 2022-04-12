package handlers

import (
	"fmt"
	"github.com/decadev/squad10/healthplus/DBConnections"
	"github.com/decadev/squad10/healthplus/models"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"html/template"
	"net/http"
	"strconv"
)

func Register(router *chi.Mux) {
	router.Get("/", Indexhandler)
	router.Get("/registerPatient", RegisterPatientHandler)
	router.Post("/registerPatient", PostRegisterPatientHandler)
	router.Get("/patientLogin", PatientLoginHandler)
	router.Get("/doctor/{ID}", Doctorhandler)

}

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

func Doctorhandler(w http.ResponseWriter, r *http.Request) {
	t, e := template.ParseFiles("pages/doctor.html")
	if e != nil {
		fmt.Println(e)
		return
	}
	ID := chi.URLParam(r, "Id")

	var x = DBConnections.ScanDoctors()
	var y = DBConnections.ScanAppoints()
	var currentDoc models.Doctor

	for i, _ := range x {
		if x[i].ID == ID {
			currentDoc = x[i]
		}
	}

	for i, _ := range y {
		if y[i].DoctorID == ID {
			currentDoc.Bookings = append(currentDoc.Bookings, y[i])
		}
	}

	e = t.Execute(w, currentDoc)
	if e != nil {
		fmt.Println(e)
		return
	}

}

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

func PostRegisterPatientHandler(w http.ResponseWriter, r *http.Request) {
	var p models.Patient
	r.ParseForm()
	name := r.FormValue("name")
	age := r.FormValue("age")
	email := r.FormValue("email")
	password := r.FormValue("password")

	g, _ := strconv.Atoi(age)
	j := uint(g)

	p.ID = uuid.NewString()
	p.Name = name
	p.Age = j
	p.Email = email
	p.Password = password

	DBConnections.InsertPatientDetails(p)
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
