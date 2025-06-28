package user

import (
	"errors"
	"fmt"
	"regexp"
)

var (
	ErrInvalidEmail = errors.New("invalid email format")
	ErrInvalidAge   = errors.New("invalid age: must be between 0 and 150")
	ErrEmptyName    = errors.New("name cannot be empty")
)

type User struct {
	Name  string
	Age   int
	Email string
}

func NewUser(name string, age int, email string) (*User, error) {
	if name == "" {
		return nil, ErrEmptyName
	}
	if age < 0 || age > 150 {
		return nil, ErrInvalidAge
	}
	if !IsValidEmail(email) {
		return nil, ErrInvalidEmail
	}
	u := &User{Name: name, Age: age, Email: email}
	return u, nil
}
func (u *User) Validate() error {
	if u.Name == "" {
		return ErrEmptyName
	}
	if u.Age < 0 || u.Age > 150 {
		return ErrInvalidAge
	}
	if !IsValidEmail(u.Email) {
		return ErrInvalidEmail
	}
	return nil
}
func (u *User) String() string {
	return fmt.Sprintf("%s (%d) - %s", u.Name, u.Age, u.Email)
}
func IsValidEmail(email string) bool {
	re := regexp.MustCompile(`^[^@\s]+@[^@\s]+\.[^@\s]+$`)
	return re.MatchString(email)
}
