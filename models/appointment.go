package models

type Appointment struct {
	ID              string `json:"id" gorm:"primaryKey"`
	Purpose         string `json:"purpose"`
	PatientID       string `json:"patientID" gorm:"size:32"`
	DoctorID        string `json:"doctorID" gorm:"size:32"`
	Date            string `json:"date"`
	AppointmentHour string `json:"appointmentHour"`
}
