package customError

type CustomError struct {
	Message string
	Err     error
	Type    string
}

func (ce CustomError) Error() string {
	return "Error with message: " + ce.Message + " and type: " + ce.Type
}
