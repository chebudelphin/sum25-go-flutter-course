package user

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var (
<<<<<<< HEAD
	ErrInvalidEmail = errors.New("invalid email format")
	ErrInvalidAge   = errors.New("invalid age: must be between 0 and 150")
	ErrEmptyName    = errors.New("name cannot be empty")
=======
	ErrInvalidName  = errors.New("invalid name: must be between 1 and 30 characters")
	ErrInvalidAge   = errors.New("invalid age: must be between 0 and 150")
	ErrInvalidEmail = errors.New("invalid email format")
>>>>>>> e6d76f7 (update lab1 and workflow of submission)
)

type User struct {
	Name  string
	Age   int
	Email string
}

/* ---------- Валидация ---------- */

func IsValidName(name string) bool {
	name = strings.TrimSpace(name)
	return len(name) > 0 && len(name) <= 30
}

func IsValidAge(age int) bool {
	return age >= 0 && age <= 150
}

var emailRE = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

func IsValidEmail(email string) bool { return emailRE.MatchString(email) }

/* ---------- Методы ---------- */

func (u *User) Validate() error {
	switch {
	case !IsValidName(u.Name):
		return ErrInvalidName
	case !IsValidAge(u.Age):
		return ErrInvalidAge
	case !IsValidEmail(u.Email):
		return ErrInvalidEmail
	default:
		return nil
	}
}

func (u *User) String() string {
	return fmt.Sprintf("Name: %s, Age: %d, Email: %s", u.Name, u.Age, u.Email)
}

/* ---------- Фабрика ---------- */

func NewUser(name string, age int, email string) (*User, error) {
	u := &User{Name: name, Age: age, Email: email}
	if err := u.Validate(); err != nil {
		return nil, err
	}
	return u, nil
}
