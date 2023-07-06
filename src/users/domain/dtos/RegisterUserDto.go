package users

import (
	shared "ecommerce/shared/domain"
	"net/mail"
)

type RegistreUserDto struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (dto RegistreUserDto) Validate() []error {
	var errorsList []error = nil

	if dto.Username == "" {
		errorsList = append(errorsList, shared.NewDomainError("product name can't be empty"))
	}

	if dto.Email == "" {
		errorsList = append(errorsList, shared.NewDomainError("product name can't be empty"))
	}

	_, err := mail.ParseAddress(dto.Email)
	if err != nil {
		errorsList = append(errorsList, shared.NewDomainError("invalid email"))
	}

	if dto.Password == "" {
		errorsList = append(errorsList, shared.NewDomainError("product name can't be empty"))
	}

	return errorsList
}
