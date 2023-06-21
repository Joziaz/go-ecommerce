package shared

import (
	"sync"

	domErrors "ecommerce/shared/domain"
	shared "ecommerce/shared/domain/interfaces"

	"github.com/google/uuid"
)

var db map[uuid.UUID]shared.Entity = make(map[uuid.UUID]shared.Entity)

type MemoryRepository[T shared.Entity] struct {
	mutex sync.Mutex
}

func (repo *MemoryRepository[T]) Save(entity T) T {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	entity.SetId(uuid.New())
	db[entity.GetId()] = entity

	return entity
}

func (repo *MemoryRepository[T]) GetById(id uuid.UUID) (T, error) {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	entity, exist := db[id].(T)
	if !exist {
		return entity, domErrors.ErrNotFound
	}

	return entity, nil
}

func (repo *MemoryRepository[T]) GetAll() []T {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()
	var entities []T

	for _, entity := range db {
		entities = append(entities, entity.(T))
	}

	return entities
}

func (repo *MemoryRepository[T]) Update(entity T) error {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	id := entity.GetId()
	_, exist := db[id]
	if !exist {
		return domErrors.ErrNotFound
	}

	return nil
}

func (repo *MemoryRepository[T]) Delete(id uuid.UUID) error {
	_, err := repo.GetById(id)
	if err != nil {
		return err
	}

	repo.mutex.Lock()
	defer repo.mutex.Unlock()
	delete(db, id)
	return nil
}
