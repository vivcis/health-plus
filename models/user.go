package models

type User struct {
	ID           string `gorm:"primaryKey" json:"id"`
	Name         string `json:"name"`
	Age          uint   `json:"age"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
	Password     string `json:"password,omitempty"`
}
