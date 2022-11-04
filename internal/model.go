package internal

import (
	"fmt"
	_model "gin-code-generator/internal/pkg/model"
	"gin-code-generator/internal/pkg/util"
	"strings"
	"time"
)

type Model struct {
	Base
	dbtype string
}

type ModelTemplateData struct {
	TemplateData
	TableName string
	Fields    string
}

func NewModel(config Config, dbtype string) *Model {
	return &Model{
		Base{
			Date:         time.Now().Format("2006/01/02 15:04"),
			TargetPath:   "/model/",
			TemplateFile: "assets/code_template/model.tpl",
			FileSuffix:   "Model.go",
			Config:       config,
		},
		dbtype,
	}
}

func NewRequestModel(config Config, dbtype string) *Model {
	return &Model{
		Base{
			Date:         time.Now().Format("2006/01/02 15:04"),
			TargetPath:   "/model/request/",
			TemplateFile: "assets/code_template/request_model.tpl",
			FileSuffix:   "Model.go",
			Config:       config,
		},
		dbtype,
	}
}

func NewResponseModel(config Config, dbtype string) *Model {
	return &Model{
		Base{
			Date:         time.Now().Format("2006/01/02 15:04"),
			TargetPath:   "/model/response/",
			TemplateFile: "assets/code_template/response_model.tpl",
			FileSuffix:   "Model.go",
			Config:       config,
		},
		dbtype,
	}
}

func (model *Model) BuildContent() (string, error) {
	var tableName string
	switch model.dbtype {
	case "oracle":
		tableName = strings.ToUpper(model.Config.Name)
	default:
		tableName = model.Config.Name
	}

	fields, err := model.BuildFieldString(tableName)
	if err != nil {
		return "", err
	}

	data := &ModelTemplateData{
		TemplateData: TemplateData{
			Author: model.Config.Author,
			Date:   model.Date,
			Name:   model.BuildName(),
		},
		TableName: tableName,
		Fields:    fields,
	}

	return util.ParseTemplateFromAssets(model.TemplateFile, data)
}

func (model *Model) BuildFieldString(tableName string) (string, error) {
	fmt.Println(model.dbtype)
	switch model.dbtype {
	case "oracle":
		return _model.NewOracleModelFieldsBuilder(tableName).Create()
	case "postgreSQL":
		return _model.NewPostgresModelColumnsBuilder(tableName).Create()
	default:
		return _model.NewMysqlModelColumnsBuilder("", tableName).Create()
	}
}

func (model *Model) Gen() (bool, error) {
	file := model.GetTarget()

	content, err := model.BuildContent()

	if err != nil {
		return false, err
	}

	return model.Write(file, content)
}
