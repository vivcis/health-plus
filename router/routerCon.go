package router

import (
	"fmt"
	"github.com/decadev/squad10/healthplus/handlers"
	"github.com/go-chi/chi"
	"net/http"
)

func SetupRouter() {
	router := chi.NewRouter()
	fmt.Println("Server up and Running")

	router.Get("/", handlers.Indexhandler)
	router.Get("/registerPatient", handlers.RegisterPatientHandler)
	router.Post("/registerPatient", handlers.PostRegisterPatientHandler)
	router.Get("/patientLogin", handlers.PatientLoginHandler)
	//router.Get("/doctor/{ID}", Doctorhandler)

	e := http.ListenAndServe(":8084", router)
	if e != nil {
		fmt.Println(e)
		return
	}

}

//func Doctorhandler(w http.ResponseWriter, r *http.Request) {
//	t, e := template.ParseFiles("pages/doctor.html")
//	if e != nil {
//		fmt.Println(e)
//		return
//	}
//	ID := chi.URLParam(r, "Id")
//
//	var x = ScanDoctors()
//	var y = ScanAppoints()
//	var currentDoc models.Doctor
//
//	for i, _ := range x {
//		if x[i].ID == ID {
//			currentDoc = x[i]
//		}
//	}
//
//	for i, _ := range y {
//		if y[i].DoctorID == ID {
//			currentDoc.Bookings = append(currentDoc.Bookings, y[i])
//		}
//	}
//
//	e = t.Execute(w, currentDoc)
//	if e != nil {
//		fmt.Println(e)
//		return
//	}
//
//}
