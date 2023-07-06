package users

import shared "ecommerce/shared/domain/models"

type User struct {
	shared.BaseEntity
	Username string
	Email    string
	Password string
	IsActive bool
}

func NewUser(userName, email, password string) User {
	return User{
		Username: userName,
		Email:    email,
		Password: password,
	}
}
