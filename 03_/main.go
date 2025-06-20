package main

import (
	"errors"
	"fmt"
	"log"
	"ron/go/customError"
	"strconv"
)

func main() {

	// parsing age
	_, cerr := ParseAge("fdsaz")
	if cerr.Message != "" {
		log.Println("Error: ", cerr.Message, "Error Type: ", cerr.Type)
	}

	_, cerr = ParseAge("130")
	if cerr.Message != "" {
		log.Println("Error:", cerr.Message, "Error Type:", cerr.Type)
	}
	age, _ := ParseAge("30")
	if cerr.Message != "" {
		log.Println("Error: ", cerr.Message, "Error Type: ", cerr.Type)
	}
	fmt.Println(*age)

	// user1 := users.NewUser("John Doe", 30, "7m5K7@example.com", "123 Main St", true)
	// println(user1.CapitalizedUserName())

}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return float64(a) / float64(b), nil
}

func ParseAge(age string) (*int, customError.CustomError) {
	convertedAge, err := strconv.Atoi(age)
	if err != nil {
		return nil, customError.CustomError{

			Message: "Invalid age",
			Err:     err,
			Type:    "ParseError",
		}
	}
	if convertedAge < 0 || convertedAge > 120 {
		return nil, customError.CustomError{
			Message: "Invalid age range",
			Err:     err,
			Type:    "range: (0-120)"}
	}

	return &convertedAge, customError.CustomError{}
}
