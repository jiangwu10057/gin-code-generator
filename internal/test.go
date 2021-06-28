package internal

import (
	"gin-code-generator/internal/pkg/util"
	"strings"
	"time"
)

type Test struct {
	Base
	Module string
}

type TestTemplateData struct {
	TemplateData
	Package  string
	Function []string
	//可以用 gotests工具自动生成
}

func NewTest(config Config, module string) *Test {
	return &Test{
		Base{
			Date:         time.Now().Format("2006/01/02 15:04"),
			TargetPath:   "/",
			TemplateFile: "assets/code_template/test.tpl",
			FileSuffix:   "_test.go",
			Config:       config,
		},
		module,
	}
}

func (test *Test) BuildContent() (string, error) {
	return util.ParseTemplateFromAssets(test.TemplateFile, &TestTemplateData{
		TemplateData: TemplateData{
			Author: test.Config.Author,
			Date:   test.Date,
			Name:   test.BuildName(),
		},
		Package: test.Module,
	})
}

func (test *Test) GetTarget() string {
	return test.Config.Path + "/" + test.Module + test.TargetPath + test.BuildName() + strings.Title(test.Module) + test.FileSuffix
}

func (base *Test) Gen() (bool, error) {
	file := base.GetTarget()

	content, err1 := base.BuildContent()

	if err1 != nil {
		return false, err1
	}

	return base.Write(file, content)
}
