package internal

import (
	"errors"
	"gin-code-generator/internal/pkg/util"
	"io"
	"os"
	"strings"
	"time"
)

type Config struct {
	Module   string
	Name     string
	Path     string
	Author   string
	WithTest bool
}

type Generator interface {
	Init() (bool, error)
	BuildName() string
	GetTarget() string
	BuildContent() (string, error)
	Write(file string, content string) (bool, error)
	Gen() (bool, error)
}

type Base struct {
	Date         string
	TargetPath   string
	TemplateFile string
	FileSuffix   string
	Config       Config
}

type Project struct {
	Base
}

func NewProject(config Config) Project {
	return Project{
		Base{
			Date:         time.Now().Format("2006/01/02 15:04"),
			TargetPath:   "/",
			TemplateFile: "assets/",
			FileSuffix:   "",
			Config:       config,
		},
	}
}

type Service struct {
	Base
}

func NewService(config Config) *Service {
	return &Service{
		Base{
			Date:         time.Now().Format("2006/01/02 15:04"),
			TargetPath:   "/service/",
			TemplateFile: "assets/code_template/service.tpl",
			FileSuffix:   "Service.go",
			Config:       config,
		},
	}
}

type TemplateData struct {
	Author string
	Date   string
	Name   string
}

// func NewTemplateData(base interface{}) TemplateData {
// 	return TemplateData{
// 		Author: base.Config.Author,
// 		Date:   base.Date,
// 		Name:   base.BuildName(),
// 	}
// }

type ServiceTemplateData struct {
	TemplateData
}

type Router struct {
	Base
}

func NewRouter(config Config) *Router {
	return &Router{
		Base{
			Date:         time.Now().Format("2006/01/02 15:04"),
			TargetPath:   "/router/",
			TemplateFile: "assets/code_template/router.tpl",
			FileSuffix:   "Router.go",
			Config:       config,
		},
	}
}

type RouterTemplateData struct {
	TemplateData
}

type Model struct {
	Base
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

type ModelTemplateData struct {
	TemplateData
	TableName string
	//struct 可以用https://github.com/Shelnutt2/db2struct生成
}

type Test struct {
	Base
	Module string
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

// func NewQuick(config Config) {
// var newMap []Base{}
// for key, value := range []string{"service","model","router"} {
// 	switch value:
// 		case "service":{
// 			newMap = append(newMap, NewService(config))
// 		},
// 		case "model":{
// 			newMap = append(newMap, NewModel(config))
// 		},
// 		case "router":{
// 			newMap = append(newMap, NewRouter(config))
// 		}
// }

// return newMap
// }

type TestTemplateData struct {
	TemplateData
	Package  string
	Function []string
	//可以用 gotests工具自动生成
}

func (base *Base) Init() (bool, error) {
	if base.Config.Name == "" {
		return false, errors.New("name cannot be empty")
	}

	return true, nil
}

func (base *Base) BuildName() string {
	return util.Case2Camel(strings.ToLower(base.Config.Name))
}

func (base *Base) BuildContent() (string, error) {
	return util.ParseTemplateFromAssets(base.TemplateFile, nil)
}

func (service *Service) BuildContent() (string, error) {
	return util.ParseTemplateFromAssets(service.TemplateFile, &ServiceTemplateData{
		TemplateData{
			Author: service.Config.Author,
			Date:   service.Date,
			Name:   service.BuildName(),
		},
	})
}

func (model *Model) BuildContent() (string, error) {
	return util.ParseTemplateFromAssets(model.TemplateFile, &ModelTemplateData{
		TemplateData: TemplateData{
			Author: model.Config.Author,
			Date:   model.Date,
			Name:   model.BuildName(),
		},
		TableName: strings.ToUpper(model.Config.Name),
	})
}

func (router *Router) BuildContent() (string, error) {
	return util.ParseTemplateFromAssets(router.TemplateFile, &RouterTemplateData{
		TemplateData{
			Author: router.Config.Author,
			Date:   router.Date,
			Name:   router.BuildName(),
		},
	})
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

func (base *Base) GetTarget() string {
	return base.Config.Path + base.TargetPath + base.BuildName() + base.FileSuffix
}

func (test *Test) GetTarget() string {
	return test.Config.Path + "/" + test.Module + test.TargetPath + test.BuildName() + strings.Title(test.Module) + test.FileSuffix
}

func (base *Base) Write(file string, content string) (bool, error) {

	if util.CheckFileIsExist(file) {
		return true, errors.New(file + " already exists")
	}

	f, err := os.Create(file)

	if err != nil {
		return false, err
	}

	_, err1 := io.WriteString(f, content)
	return err1 == nil, err1
}

func (base *Base) Gen() (bool, error) {
	file := base.GetTarget()

	content, err1 := base.BuildContent()

	if err1 != nil {
		return false, err1
	}

	return base.Write(file, content)
}

func (base *Service) Gen() (bool, error) {
	file := base.GetTarget()

	content, err1 := base.BuildContent()

	if err1 != nil {
		return false, err1
	}

	return base.Write(file, content)
}

func (base *Router) Gen() (bool, error) {
	file := base.GetTarget()

	content, err1 := base.BuildContent()

	if err1 != nil {
		return false, err1
	}

	return base.Write(file, content)
}

func (base *Model) Gen() (bool, error) {
	file := base.GetTarget()

	content, err1 := base.BuildContent()

	if err1 != nil {
		return false, err1
	}

	return base.Write(file, content)
}

func (base *Test) Gen() (bool, error) {
	file := base.GetTarget()

	content, err1 := base.BuildContent()

	if err1 != nil {
		return false, err1
	}

	return base.Write(file, content)
}
