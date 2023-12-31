package shared

import "github.com/google/uuid"

type Repository[T Entity] interface {
	Save(T) (T, error)
	GetById(uuid.UUID) (T, error)
	GetAll() []T
	Update(T) error
	Delete(uuid.UUID) error
}
