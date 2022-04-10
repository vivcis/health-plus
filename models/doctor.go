package models

type Doctor struct {
	User
	Specialty   string `json:"specialty"`
	WorkingHour string `json:"workingHour"`
	Bookings    []Appointment
}
