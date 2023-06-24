package shared

type DomainError struct {
	msg string
}

func (e *DomainError) Error() string {
	return e.msg
}

var (
	ErrNotFound      = &DomainError{"not found"}
	ErrDuplicatedKey = &DomainError{"duplicated key not allowed"}
)
