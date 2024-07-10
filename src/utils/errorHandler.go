package utils

import "encoding/json"

type Err struct {
	Code    int
	Message string
	Detail  string
}

// Error implements error.
func (e *Err) Error() string {
	err := Err{
		Code:    e.Code,
		Message: e.Message,
		Detail:  e.Detail,
	}
	b, _ := json.Marshal(err)

	return string(b)
}

func (e *Err) format()

func NewError(code int, message, detail string) error {
	return &Err{
		Code:    code,
		Message: message,
		Detail:  detail,
	}
}

//Ensure error type implements error interface
var _ error = &Err{}
