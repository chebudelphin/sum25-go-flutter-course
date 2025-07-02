package user

import (
	"errors"
	"fmt"
	"regexp"
)

// Predefined errors
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

<<<<<<< HEAD
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
=======
// Validate checks if the user data is valid, returns an error for each invalid field
func (u *User) Validate() error {
	if !IsValidName(u.Name) {
		return ErrInvalidName
	}

	if !IsValidAge(u.Age) {
		return ErrInvalidAge
	}

	if !IsValidEmail(u.Email) {
		return ErrInvalidEmail
	}

	return nil
}

// String returns a string representation of the user, formatted as "Name: <name>, Age: <age>, Email: <email>"
func (u *User) String() string {
	// TODO: Implement this function
	return ""
}

// NewUser creates a new user with validation, returns an error if the user is not valid
func NewUser(name string, age int, email string) (*User, error) {
	// TODO: Implement this function
	return nil, nil
}

// IsValidEmail checks if the email format is valid
// You can use regexp.MustCompile to compile the email regex
func IsValidEmail(email string) bool {
	// TODO: Implement this function
	return false
>>>>>>> e6d76f7 (update lab1 and workflow of submission)
}

// IsValidName checks if the name is valid, returns false if the name is empty or longer than 30 characters
func IsValidName(name string) bool {
	// TODO: Implement this function
	return false
}

// IsValidAge checks if the age is valid, returns false if the age is not between 0 and 150
func IsValidAge(age int) bool {
	// TODO: Implement this function
	return false
}