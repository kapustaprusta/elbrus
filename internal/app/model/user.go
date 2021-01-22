package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

// User ...
type User struct {
	ID               int
	Name             string
	Password         string
	EncryptedPasword string
}

// Validate ...
func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Name, validation.Required, is.ASCII),
		validation.Field(&u.Password, validation.By(requiredIf(u.EncryptedPasword == "")), validation.Length(6, 100)),
	)
}

// BeforeCreate ...
func (u *User) BeforeCreate() error {
	if (len(u.Password)) > 0 {
		enc, err := encryptString(u.Password)
		if err != nil {
			return err
		}

		u.EncryptedPasword = enc
	}

	return nil
}

func encryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
