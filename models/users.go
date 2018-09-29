package models

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username 	string `json:"username" gorm:"not null"`
	Email		string `json:"email" gorm:"not null;unique"`
}

// alias name of struct
type PostItem struct {
	gorm.Model
	Title	string `json:"title" gorm:"not null"`
}

func (item PostItem) TableName() string {
	return "posting"
}

func (u User) ValidationUser() error {
	if len(strings.Trim(u.Username, " ")) < 1 { // checked Trim
		return fmt.Errorf("masukkan username")
	}

	if len(strings.Trim(u.Email, " ")) < 1 {
		return fmt.Errorf("masukkan email")
	}

	return nil
}

type Activities struct {
	ID	int `json:"id" gorm:"id"`
	Running	int `json:"running" gorm:"running"`
	Walking	int `json:"walking" gorm:"walking"`
}

type UserList struct {
	Data []struct {
		ID	int    `json:"id"`
		Name	string `json:"name"`
		Address	string `json:"address"`
	} `json:"data"`
	Message	string `json:"message"`
}

type UserJWT struct {
	UserName	string
	FirstName	string
	LastName	string
}