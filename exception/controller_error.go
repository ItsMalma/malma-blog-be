package exception

import (
	"fmt"
)

type ControllerError struct {
	Msg  string
	Err  string
	Code int
}

func (err ControllerError) Error() string {
	return fmt.Sprintf("%v %v: %v", err.Msg, err.Code, err.Err)
}

func ErrParseRequest(reason string) ControllerError {
	return ControllerError{
		Msg:  "ERR_PARSE_REQUEST",
		Err:  fmt.Sprintf("Unable to parse request (%v)", reason),
		Code: 422,
	}
}

func ErrContentType(expected string, got string) ControllerError {
	return ControllerError{
		Msg:  "ERR_UNSUPPORTED_CONTENT_TYPE",
		Err:  fmt.Sprintf("Expected %v content type but got %v instead", expected, got),
		Code: 415,
	}
}
