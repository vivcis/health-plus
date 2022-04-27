package models

import (
	"strings"
)

type User struct {
	ID           string `gorm:"primaryKey" json:"id"`
	Username     string `gorm:"unique" json:"username"`
	Name         string `json:"name"`
	Age          uint   `json:"age"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
	Password     string `json:"password,omitempty" gorm:"-"`
}

func Capitalise(s string) string {
	s = strings.ToLower(s)
	var newString []byte
	for i := 0; i < len(s); i++ {
		if i == 0 {
			newString = append(newString, s[i]-32)
			continue
		}
		newString = append(newString, s[i])
	}
	return string(newString)
}
