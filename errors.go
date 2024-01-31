package templateless

import "fmt"

type BadRequestCode int

const (
	AccountQuotaExceeded BadRequestCode = 200
	ProviderKeyMissing   BadRequestCode = 300
	ProviderKeyInvalid   BadRequestCode = 301
)

func (code BadRequestCode) Error() string {
	switch code {
	case AccountQuotaExceeded:
		return "account quota exceeded"
	case ProviderKeyMissing:
		return "provider key missing"
	case ProviderKeyInvalid:
		return "provider key invalid"
	default:
		return "unknown bad request code"
	}
}

type CustomError struct {
	Type           string         `json:"type"`
	BadRequestCode BadRequestCode `json:"code"`
	Field          string         `json:"field"`
	ErrorMsg       string         `json:"error"`
}

func (e *CustomError) Error() string {
	return e.ErrorMsg
}

func UnauthorizedError() *CustomError {
	return &CustomError{Type: "Unauthorized", ErrorMsg: "unauthorized"}
}

func ForbiddenError() *CustomError {
	return &CustomError{Type: "Forbidden", ErrorMsg: "forbidden"}
}

func InvalidParameterError(field string) *CustomError {
	return &CustomError{
		Type:     "InvalidParameter",
		Field:    field,
		ErrorMsg: fmt.Sprintf("invalid parameter: %s", field),
	}
}

func BadRequestError(code BadRequestCode, error string) *CustomError {
	return &CustomError{
		Type:           "BadRequest",
		BadRequestCode: code,
		ErrorMsg:       fmt.Sprintf("bad request: %s - %s", code.Error(), error),
	}
}

func UnavailableError() *CustomError {
	return &CustomError{
		Type:     "Unavailable",
		ErrorMsg: "unavailable",
	}
}

func UnknownError() *CustomError {
	return &CustomError{
		Type:     "UnknownError",
		ErrorMsg: "unknown error",
	}
}
