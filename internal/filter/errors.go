package filter

import "fmt"

type Error struct {
	Err string
}

func (e *Error) Error() string {
	return e.Err
}

func NewError(err string) *Error {
	return &Error{
		Err: err,
	}
}

func ErrInvalidBandwidthLimit(msg string, limit string) *Error {
	return NewError(fmt.Sprintf("Invalid bandwidth limit: %s, %s", limit, msg))
}

func ErrInvalidPacketsLimit(msg string, limit string) *Error {
	return NewError(fmt.Sprintf("Invalid packets limit: %s, %s", limit, msg))
}

func ErrInvalidIP(msg string, ip string) *Error {
	return NewError(fmt.Sprintf("Invalid IP: %s, %s", ip, msg))
}

func ErrInvalidDev(msg string, dev string) *Error {
	return NewError(fmt.Sprintf("Invalid network interface: %s, %s", dev, msg))
}

func ErrWrongOS(os string) *Error {
	return NewError(fmt.Sprintf("OS %s is not supported", os))
}
