package shared

type EntityDB[T Entity] interface {
	ToEntity() T
	FromEntity(T) any
	TableName() string
}
