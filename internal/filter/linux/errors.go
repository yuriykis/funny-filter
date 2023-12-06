package linux

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

func ErrInvalidLinuxConfigType() *Error {
	return NewError("Invalid Linux config type")
}
