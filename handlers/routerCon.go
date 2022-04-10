package handlers

import (
	"fmt"
	"github.com/decadev/squad10/healthplus/DBConnections"
	"github.com/decadev/squad10/healthplus/models"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
)

func Register(router *chi.Mux) {
	router.Get("/", Indexhandler)
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
