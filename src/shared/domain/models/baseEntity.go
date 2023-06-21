package shared

import (
	"time"

	"github.com/google/uuid"
)

type BaseEntity struct {
	id         uuid.UUID
	CreatedAt  time.Time `json:"createdAt"`
	LastUpdate time.Time `json:"LastUpdate"`
}

func (entity BaseEntity) GetId() uuid.UUID {
	return entity.id
}

func (entity *BaseEntity) SetId(id uuid.UUID) {
	entity.id = id
}
