package calculator

import (
	"errors"
	"strconv"
)

var ErrDivisionByZero = errors.New("division by zero")
var ErrInvalidNumber = errors.New("invalid number format")

func Add(a, b float64) float64 {
	return a + b
}
func Subtract(a, b float64) float64 {
	return a - b
}
func Multiply(a, b float64) float64 {
	return a * b
}
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivisionByZero
	}
	return a / b, nil
}
func StringToFloat(s string) (float64, error) {
	if s == "" {
		return 0, ErrInvalidNumber
	}
	val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, ErrInvalidNumber
	}
	return val, nil
}
func FloatToString(f float64, precision int) string {
	return strconv.FormatFloat(f, 'f', precision, 64)
}
