package shared

import (
	domainErrors "ecommerce/shared/domain"
	shared "ecommerce/shared/domain/interfaces"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormRepository[T shared.Entity, K shared.EntityDB[T]] struct {
	db *gorm.DB
}

func NewGormRepository[T shared.Entity, K shared.EntityDB[T]](db *gorm.DB) GormRepository[T, K] {
	return GormRepository[T, K]{db}
}

func (repo GormRepository[T, K]) GetAll() []T {
	var dbEntities []K
	repo.db.Find(&dbEntities)

	var entities = make([]T, len(dbEntities))
	for i, v := range dbEntities {
		entities[i] = v.ToEntity()
	}
	return entities
}

func (repo GormRepository[T, K]) GetById(id uuid.UUID) (T, error) {
	var entity K
	result := repo.db.First(&entity, "id = ?", id.String())

	if result.Error != nil {
		return entity.ToEntity(), result.Error
	}

	return entity.ToEntity(), nil
}

func (repo GormRepository[T, K]) Save(entity T) (T, error) {
	var entityDb K
	entityDb = entityDb.FromEntity(entity).(K)

	result := repo.db.Create(entityDb)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return entity, domainErrors.ErrDuplicatedKey
		}

		return entity, result.Error
	}

	return entity, nil
}

func (repo GormRepository[T, K]) Update(entity T) error {
	var entityDb K
	entityDb = entityDb.FromEntity(entity).(K)

	result := repo.db.Updates(entityDb)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo GormRepository[T, K]) Delete(id uuid.UUID) error {
	var entity K
	result := repo.db.Delete(&entity, id.String())
	if result.Error != nil {
		return result.Error
	}

	return nil
}
