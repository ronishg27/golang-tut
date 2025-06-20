package day_06

import "fmt"

type CustomError struct {
	message    string
	errContext string
	lineNumber int
	fileName   string
}

func (c CustomError) Error() string {
	return fmt.Sprintln("Message:", c.message, "Error Context:", c.errContext, "Line Number:", c.lineNumber, "File Name:", c.fileName)
}
