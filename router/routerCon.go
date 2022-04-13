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

	router.Get("/registerDoctor", handlers.RegisterDoctorHandler)
	router.Post("/registerDoctor", handlers.PostRegisterDoctorHandler)
	router.Get("/doctorLogin", handlers.DoctorLoginHandler)

	e := http.ListenAndServe(":8084", router)
	if e != nil {
		fmt.Println(e)
		return
	}

}
