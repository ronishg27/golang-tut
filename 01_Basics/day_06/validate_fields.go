package day_06

import (
	"errors"
	"log"
	"math"
	"strings"
)

func ValidateFields(value string, fieldName string) error {

	ce := &CustomError{}

	if fieldName == "address" {

		if len(value) < 5 {
			ce.errContext = "Validating AddressError"
			ce.message = "Address must be at least 5 characters long"
			ce.lineNumber = 7
			ce.fileName = "validate_fields.go"
			return ce
		}
	}
	if fieldName == "name" {
		if len(value) < 5 {
			ce.errContext = "Validating Name Error"
			ce.message = "Name must be at least 5 characters long"
			ce.lineNumber = 7
			ce.fileName = "validate_fields.go"
			return ce
		}
	}

	return nil
}

func CheckError() {
	name := "a"
	address := "a"
	err := ValidateFields(name, "name")
	log.Println("Error:", err.Error())

	err = ValidateFields(address, "address")
	log.Println("Error:", err.Error())

}

func SquareRoot(x float64) (float64, error) {
	if x < 0 {
		return -1, errors.New("negative number detected")
	}
	sqrt := math.Sqrt(x)
	return sqrt, nil

}

func CheckName(name string) (string, error) {
	if len(name) < 5 {
		return "", errors.New("name is too short")
	}
	return strings.ToUpper(name), nil
}
