package model

import (
	"gin-code-generator/internal/pkg/util"
	"strings"
)

type PostgresColumn struct {
	Comment    string `gorm:"column:description"`
	ColumnName string `gorm:"column:attname"`
	DataType   string `gorm:"column:data_type"`
}

func (PostgresColumn) TableName() string {
	return "pg_catalog.pg_attribute"
}

type PostgresColumnGetter struct {
	Result []PostgresColumn
}

func NewPostgresColumnGetter() *PostgresColumnGetter {
	return &PostgresColumnGetter{}
}

func (getter *PostgresColumnGetter) Get(tableName string) error {
	err := DB.Raw(`SELECT 
	( SELECT description FROM pg_catalog.pg_description WHERE objoid = A.attrelid AND objsubid = A.attnum ) AS description,
	A.attname,
	( select typname from pg_type where oid = A.atttypid) AS data_type
FROM
	pg_catalog.pg_attribute A
WHERE
	1 = 1
	AND A.attrelid = ( SELECT oid FROM pg_class WHERE relname = ? )
	AND A.attnum > 0
	AND NOT A.attisdropped
ORDER BY
	A.attnum;`, tableName).Scan(&getter.Result).Error
	return err
}

type PostgresModelColumnsBuilder struct {
	TableName string
}

func NewPostgresModelColumnsBuilder(tableName string) *PostgresModelColumnsBuilder {
	return &PostgresModelColumnsBuilder{
		TableName: tableName,
	}
}

func (builder *PostgresModelColumnsBuilder) GetColumns() ([]PostgresColumn, error) {
	getter := NewPostgresColumnGetter()
	err := getter.Get(builder.TableName)
	if err != nil {
		return []PostgresColumn{}, err
	}

	return getter.Result, nil
}

func (builder *PostgresModelColumnsBuilder) CoverColumnType(t string) string {

	switch t {
	case "varchar", "char", "text", "longtext", "mediumtext", "uuid":
		return "string"
	case "bigint", "int2", "int4", "int8":
		return "uint"
	case "date", "timestamp":
		return "time.Time"
	default:
		return t
	}
}

func (builder *PostgresModelColumnsBuilder) Create() (string, error) {
	var str strings.Builder

	columns, err := builder.GetColumns()

	if err != nil {
		return "", err
	}

	for _, column := range columns {
		name := column.ColumnName
		field := util.Case2Camel(strings.ToLower(name))

		line := "	" + field + "	" + builder.CoverColumnType(column.DataType) + " `gorm:\"column:" + name + ";comment:" + column.Comment +
			"\" json:\"" + util.LowerFirst(field) + "\"`	//" + column.Comment + "\n"

		str.WriteString(line)
	}

	return str.String(), nil

}
