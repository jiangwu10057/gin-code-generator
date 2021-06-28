package internal

import (
	_model "gin-code-generator/internal/pkg/model"
	"gin-code-generator/internal/pkg/util"
	"strings"
	"time"
)

type Model struct {
	Base
}

type ModelTemplateData struct {
	TemplateData
	TableName string
	Fields    string
}

func NewModel(config Config) *Model {
	return &Model{
		Base{
			Date:         time.Now().Format("2006/01/02 15:04"),
			TargetPath:   "/model/",
			TemplateFile: "assets/code_template/model.tpl",
			FileSuffix:   "Model.go",
			Config:       config,
		},
	}
}

func NewRequestModel(config Config) *Model {
	return &Model{
		Base{
			Date:         time.Now().Format("2006/01/02 15:04"),
			TargetPath:   "/model/request/",
			TemplateFile: "assets/code_template/request_model.tpl",
			FileSuffix:   "Model.go",
			Config:       config,
		},
	}
}

func NewResponseModel(config Config) *Model {
	return &Model{
		Base{
			Date:         time.Now().Format("2006/01/02 15:04"),
			TargetPath:   "/model/response/",
			TemplateFile: "assets/code_template/response_model.tpl",
			FileSuffix:   "Model.go",
			Config:       config,
		},
	}
}

func (model *Model) BuildContent() (string, error) {
	tableName := strings.ToUpper(model.Config.Name)
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
	builder := _model.NewOracleModelFieldsBuilder(tableName)
	return builder.Create()
}

func (model *Model) Gen() (bool, error) {
	file := model.GetTarget()

	content, err := model.BuildContent()

	if err != nil {
		return false, err
	}

	return model.Write(file, content)
}
