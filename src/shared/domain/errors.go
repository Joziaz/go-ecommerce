package shared

type DomainError struct {
	msg string
}

func NewDomainError(msg string) *DomainError {
	return &DomainError{msg}
}

func (e *DomainError) Error() string {
	return e.msg
}

var (
	ErrNotFound        = &DomainError{"not found"}
	ErrDuplicatedKey   = &DomainError{"duplicated key not allowed"}
	ErrValidationError = &DomainError{"a error occur during the validation"}
)
