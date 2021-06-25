package model

type MysqlColumn struct {
	ColumnName string `gorm:"column:COLUMN_NAME"`
	DataType string `gorm:"column:DATA_TYPE"`
	DataDefualt string `gorm:"column:COLUMN_DEFAULT"`
	Comment string `gorm:"column:COLUMN_COMMENT"`
	ColumnKey string `gorm:"column:COLUMN_KEY"` // =PRI就是主键

	Column
}

func (MysqlColumn) TableName() string {
	return "INFORMATION_SCHEMA.COLUMNS"
}

type MysqlColumnGetter struct {
	SchemaName string
	Result []MysqlColumn
}

func NewMysqlColumnGetter(schema string) *MysqlColumnGetter {
	return &MysqlColumnGetter{
		SchemaName: schema,
	}
}

func (getter *MysqlColumnGetter) Get(tableName string) (error) {
	err := DB.Find(&getter.Result, "table_schema=? AND table_name=?", getter.SchemaName, tableName).Error
	return err
}

type MysqlModelBuilder struct {
	
}