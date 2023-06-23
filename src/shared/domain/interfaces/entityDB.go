package shared

type EntityDB[T Entity] interface {
	ToEntity() T
	FromEntity(T) EntityDB[T]
	TableName() string
}
