package rerrors

import (
	"fmt"
	"net/http"
)

const (
	BadRequestErr           = http.StatusBadRequest
	UnauthorizedErr         = http.StatusUnauthorized
	ForbiddenErr            = http.StatusForbidden
	NotFoundErr             = http.StatusNotFound
	ConflictErr             = http.StatusConflict
	InternalErr             = http.StatusInternalServerError
	UnproccessibleEntityErr = http.StatusUnprocessableEntity
)

var (
	internalErrMsg = "failed to process the request at this time, please try again later."

	errMessages = map[int]string{
		BadRequestErr:           "invalid request error",
		UnauthorizedErr:         "email/phone_no or passcode invalid",
		ForbiddenErr:            "user is not authorized to access this resource",
		NotFoundErr:             "not found",
		ConflictErr:             "this email/phone_no already exists, please use a differenct email address",
		InternalErr:             internalErrMsg,
		UnproccessibleEntityErr: "invalid request error",
	}
)

func message(code int) string {
	if value, ok := errMessages[code]; ok {
		return value
	}
	return internalErrMsg
}

func detail(code int, err error) string {
	if _, ok := errMessages[code]; ok {
		if err == nil {
			return "No error detail"
		}
		return fmt.Sprintf("%v", err)
	}
	return "unknown"
}

func Format(code int, err error) error {
	return NewError(code, message(code), detail(code, err))
}
