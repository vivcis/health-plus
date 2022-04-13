package models

type User struct {
	ID           string `gorm:"primaryKey" json:"id"`
	Username     string `gorm:"unique" json:"username"`
	Name         string `json:"name"`
	Age          uint   `json:"age"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
	Password     string `json:"password,omitempty" gorm:"-"`
}
