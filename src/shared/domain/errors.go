package shared

type DomainError struct {
	msg string
}

func (e *DomainError) Error() string {
	return e.msg
}

var (
	ErrNotFound     = &DomainError{"not found"}
	ErrAlreadyExist = &DomainError{"alredy exist a item with that key"}
)
