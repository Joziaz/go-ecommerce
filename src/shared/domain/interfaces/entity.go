package shared

import "github.com/google/uuid"

type Entity interface {
	GetId() uuid.UUID
	SetId(uuid.UUID)
}
