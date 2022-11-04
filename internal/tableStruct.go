package internal

import (
	"fmt"
	"strings"
	"time"
)

type TableStruct struct {
	Model
}

func NewTableStruct(config Config) *TableStruct {
	return &TableStruct{
		Model{
			Base{
				Date:         time.Now().Format("2006/01/02 15:04"),
				TargetPath:   "/service/",
				TemplateFile: "assets/code_template/service.tpl",
				FileSuffix:   "Service.go",
				Config:       config,
			},
			"",
		},
	}
}

func (model *TableStruct) BuildContent() (string, error) {
	tableName := strings.ToUpper(model.Config.Name)
	fields, err := model.BuildFieldString(tableName)
	if err != nil {
		return "", err
	}

	return fields, nil
}

func (model *TableStruct) Gen() (bool, error) {

	content, err := model.BuildContent()

	if err != nil {
		return false, err
	}

	fmt.Printf(content)
	return true, nil
}
