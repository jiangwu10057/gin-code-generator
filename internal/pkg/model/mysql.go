package model

import (
	"gin-code-generator/internal/pkg/util"
	"strings"
)

type MysqlColumn struct {
	ColumnName  string `gorm:"column:COLUMN_NAME"`
	DataType    string `gorm:"column:DATA_TYPE"`
	DataDefualt string `gorm:"column:COLUMN_DEFAULT"`
	Comment     string `gorm:"column:COLUMN_COMMENT"`
	ColumnKey   string `gorm:"column:COLUMN_KEY"` // =PRI就是主键
}

func (MysqlColumn) TableName() string {
	return "INFORMATION_SCHEMA.COLUMNS"
}

type MysqlColumnGetter struct {
	SchemaName string
	Result     []MysqlColumn
}

func NewMysqlColumnGetter(schema string) *MysqlColumnGetter {
	return &MysqlColumnGetter{
		SchemaName: schema,
	}
}

func (getter *MysqlColumnGetter) Get(tableName string) error {
	err := DB.Table("INFORMATION_SCHEMA.COLUMNS").Find(&getter.Result, "table_schema=? AND table_name=?", getter.SchemaName, tableName).Error
	return err
}

type MysqlModelColumnsBuilder struct {
	SchemaName string
	TableName  string
}

func NewMysqlModelColumnsBuilder(schemaName string, tableName string) *MysqlModelColumnsBuilder {
	return &MysqlModelColumnsBuilder{
		SchemaName: schemaName,
		TableName:  tableName,
	}
}

func (builder *MysqlModelColumnsBuilder) GetColumns() ([]MysqlColumn, error) {
	getter := NewMysqlColumnGetter(builder.SchemaName)
	err := getter.Get(builder.TableName)
	if err != nil {
		return []MysqlColumn{}, err
	}

	return getter.Result, nil
}

func (builder *MysqlModelColumnsBuilder) CoverColumnType(t string) string {
	switch t {
	case "varchar", "char", "longtext", "mediumtext":
		return "string"
	case "bigint":
		return "uint"
	case "tinyint", "smallint":
		return "int"
	case "DATE":
		return "time.Time"
	default:
		return t
	}
}

func (builder *MysqlModelColumnsBuilder) Create() (string, error) {
	var str strings.Builder

	columns, err := builder.GetColumns()

	if err != nil {
		return "", err
	}

	for _, column := range columns {
		name := column.ColumnName
		field := util.Case2Camel(strings.ToLower(name))

		var pk string
		if column.ColumnKey == "PRI" {
			pk = "primary_key;"
		} else {
			pk = ""
		}

		line := "	" + field + "	" + builder.CoverColumnType(column.DataType) + " `gorm:\"column:" + name + ";" + pk + "comment:" + column.Comment +
			"\" json:\"" + util.LowerFirst(field) + "\"`	//" + column.Comment + "\n"

		str.WriteString(line)
	}

	return str.String(), nil

}
