package model

type Column interface {
	TableName() string
}

type ColumnGetter interface {
	Get(tableName string) error
}

type ModelBuilder interface {
	Create() (string, error)
}
