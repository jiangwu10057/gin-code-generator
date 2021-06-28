package model

import (
	"strings"

	"gin-code-generator/internal/pkg/util"
)

type OracleColumn struct {
	ColumnName  string `gorm:"column:COLUMN_NAME"`
	DataType    string `gorm:"column:DATA_TYPE"`
	DataDefualt string `gorm:"column:DATA_DEFAULT"`
	Comment     string
	IsPrimary   bool
}

func (OracleColumn) TableName() string {
	return "user_tab_columns"
}

type OracleColumnComment struct {
	ColumnName string `gorm:"column:COLUMN_NAME"`
	Comments   string `gorm:"column:COMMENTS"`
}

func (OracleColumnComment) TableName() string {
	return "user_col_comments"
}

type OracleColumnGetter struct {
	Result []OracleColumn
}

func NewOracleColumnGetter() *OracleColumnGetter {
	return &OracleColumnGetter{}
}

func (getter *OracleColumnGetter) Get(tableName string) error {
	tableName = strings.ToUpper(tableName)

	columns, err := getter.getColumns(tableName)
	if err != nil {
		return err
	}
	comments, err1 := getter.getColumnComments(tableName)

	if err1 != nil {
		return err1
	}

	pkColumn, err2 := getter.getPrimaryKeys(tableName)

	if err2 != nil {
		return err2
	}

	mapObj := map[string]string{}
	for _, comment := range comments {
		mapObj[comment.ColumnName] = comment.Comments
	}

	for index, column := range columns {
		if column.ColumnName == pkColumn {
			columns[index].IsPrimary = true
		} else {
			columns[index].IsPrimary = false
		}
		v, ok := mapObj[column.ColumnName]
		if ok {
			columns[index].Comment = v
		} else {
			columns[index].Comment = ""
		}
	}

	getter.Result = columns

	return nil
}

func (getter *OracleColumnGetter) getColumnComments(tableName string) ([]OracleColumnComment, error) {
	var comments []OracleColumnComment
	err := DB.Find(&comments, "Table_Name=?", tableName).Error
	return comments, err

}

func (getter *OracleColumnGetter) getColumns(tableName string) ([]OracleColumn, error) {
	var columns []OracleColumn
	err := DB.Find(&columns, "Table_Name=?", tableName).Error
	return columns, err
}

func (getter *OracleColumnGetter) getPrimaryKeys(tableName string) (pk string, err error) {
	err = DB.Raw("Select b.Column_Name From user_Constraints a,user_Cons_Columns b Where a.Constraint_Type = 'P' and a.Constraint_Name = b.Constraint_Name And a.Owner = b.Owner And a.table_name = b.table_name And a.table_name=?", tableName).Scan(&pk).Error
	return pk, err
}

type OracleModelBuilder struct {
	TableName string
	Columns   ColumnGetter
}

func NewOracleModelBuilder(tableName string, columns ColumnGetter) *OracleModelBuilder {
	return &OracleModelBuilder{
		Columns:   columns,
		TableName: tableName,
	}
}

func (builder *OracleModelBuilder) Create() (string, error) {
	var str strings.Builder

	str.WriteString("type " + util.Case2Camel(strings.ToLower(builder.TableName)) + "Model struct {\n")

	for _, column := range builder.Columns {
		name := column.ColumnName
		field := util.Case2Camel(strings.ToLower(name))

		var pk string
		if column.IsPrimary {
			pk = "primary_key;"
		} else {
			pk = ""
		}

		line := "	" + field + "	" + builder.coverColumnType(column.DataType) + "`gorm:\"column:" + name + ";" + pk + "comment:" + column.Comment +
			"\" json:\"" + field + "\"`	//" + column.Comment + "\n"

		str.WriteString(line)
	}
	str.WriteString("}")

	return str.String(), nil

}

func (builder *OracleModelBuilder) coverColumnType(t string) string {
	switch t {
	case "NVARCHAR2", "NCHAR", "VARCHAR", "VARCHAR2", "BLOB", "CLOB", "BFILE":
		return "string"
	case "NUMBER", "INTEGER": // 有风险 number类型有可能是浮点数
		return "int"
	case "DATE":
		return "time.Time"
	default:
		// return fmt.Fprintr()
		return ""
	}
}
