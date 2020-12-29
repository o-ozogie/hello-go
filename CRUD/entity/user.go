package entity

import (
	"gorm.io/gorm"
	"hello-go/CRUD/request"
)

type User struct {
	Email string `gorm:"primaryKey;size:50"`
	Password string `gorm:"size:255"`
	Name string `gorm:"size:45"`
}

type Users []User

func CreateUser(signUpRequest request.SignUpRequest) *gorm.DB {
	user := User{Email: signUpRequest.Email, Password: signUpRequest.Password, Name: signUpRequest.Name}

	return connection.FirstOrCreate(new(User), user)
}

func FindUserByEmail(email string) *gorm.DB {
	return connection.First(new(User), "email=?", email)
}
