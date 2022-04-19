package router

import (
	"fmt"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/decadev/squad10/healthplus/handlers"
	"github.com/go-chi/chi"
)

func SetupRouter() {
	handlers.Sessions  = scs.New()
	handlers.Sessions.Lifetime = 24 * time.Hour
	router := chi.NewRouter()
	fmt.Println("Server up and Running")

	fs := http.FileServer(http.Dir("./pages/static/"))
	router.Handle("/static/", http.StripPrefix("/static/", fs))

	router.Get("/", handlers.Indexhandler)
	router.Get("/registerPatient", handlers.RegisterPatientHandler)
	router.Post("/registerPatient", handlers.PostRegisterPatientHandler)
	router.Get("/patientLogin", handlers.PatientLoginHandler)

	router.Get("/registerDoctor", handlers.RegisterDoctorHandler)
	router.Post("/registerDoctor", handlers.PostRegisterDoctorHandler)
	router.Get("/doctorLogin", handlers.DoctorLoginHandler)

	router.Post("/doctorLogin", handlers.PostLoginDoctordHandler)
	router.Get("/doctorLogout", handlers.DoctorLogoutHandler)
	router.Get("/doctorDashboard", handlers.DoctorHomeHandler)

	//router.Get("/doctorHome", handlers.DoctorHomeHandler)

	e := http.ListenAndServe(":8084", handlers.Sessions.LoadAndSave(router))
	if e != nil {
		fmt.Println(e)
		return
	}

}
