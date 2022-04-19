package router

import (
	"fmt"
	"github.com/decadev/squad10/healthplus/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func SetupRouter() {
	router := mux.NewRouter()
	router.HandleFunc("/", handlers.Indexhandler).Methods("GET")
	router.HandleFunc("/registerPatient", handlers.RegisterPatientHandler).Methods("GET")
	router.HandleFunc("/registerPatient", handlers.PostRegisterPatientHandler).Methods("POST")
	router.HandleFunc("/patientLogin", handlers.PatientLoginHandler).Methods("GET")
	router.HandleFunc("/registerDoctor", handlers.RegisterDoctorHandler).Methods("GET")
	router.HandleFunc("/registerDoctor", handlers.PostRegisterDoctorHandler).Methods("POST")
	router.HandleFunc("/doctorLogin", handlers.DoctorLoginHandler).Methods("GET")

	fs := http.FileServer(http.Dir("./pages/static/"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	//http.Handle("/", router)

	fmt.Println("Server up and Running")

	port := os.Getenv("DB_PORT")

	e := http.ListenAndServe(port, router)
	if e != nil {
		fmt.Println(e)
		return
	}
}
