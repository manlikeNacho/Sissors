package rerrors

import (
	"encoding/json"
)

// Internal error type
type Err struct {
	Code    int
	Message string
	Detail  string
}

// Ensure Customized error type implements error interface
var _ error = &Err{}

// NewError returns a new customized error type
func NewError(code int, message, detail string) error {
	return &Err{
		Code:    code,
		Message: message,
		Detail:  detail,
	}
}

// Error returns a json string representation of our customized error type
// this is the only method required to implement the error interface
func (e *Err) Error() string {
	err := Err{
		Code:    e.Code,
		Message: e.Message,
		Detail:  e.Detail,
	}
	b, _ := json.Marshal(err)

	return string(b)
}

// func NewErrFromJSON(errString string) (*Err, error) {
// 	var er Err
// 	err := json.Unmarshal([]byte(errString), &er)
// 	if err != nil {
// 		return nil, err
// 	}
// 	fmt.Println(errString, err, er)
// 	return &er, nil
// }

func (e *Err) Status() int {
	return e.Code
}

func ConvertErr(err error) *Err {
	return err.(*Err)
}
