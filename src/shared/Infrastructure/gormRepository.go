package shared

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormRepository[T gorm.Model] struct {
	db gorm.DB
}

func (repo GormRepository[T]) GetAll() []T {
	var entities []T
	repo.db.Find(&entities)

	return entities
}

func (repo GormRepository[T]) GetById(id uuid.UUID) (T, error) {
	var entity T
	result := repo.db.First(&entity, "id = ?", id.String())

	if result.Error != nil {
		return entity, result.Error
	}

	return entity, nil
}

func (repo GormRepository[T]) Save(entity T) T {
	repo.db.Create(entity)

	return entity
}

func (repo GormRepository[T]) Update(entity T) error {
	repo.db.Updates()
}
