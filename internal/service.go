package internal

import (
	"gin-code-generator/internal/pkg/util"
	"time"
)

type Service struct {
	Base
}

type ServiceTemplateData struct {
	TemplateData
	NameSpace string
	WithCurd  bool
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

func (service *Service) BuildContent() (string, error) {
	namespace, err := util.GetNameSpace(service.Config.Path + "/")
	if err != nil {
		return "", err
	}
	return util.ParseTemplateFromAssets(service.TemplateFile, &ServiceTemplateData{
		TemplateData: TemplateData{
			Author: service.Config.Author,
			Date:   service.Date,
			Name:   service.BuildName(),
		},
		WithCurd:  service.Config.WithCurd,
		NameSpace: namespace,
	})
}

func (service *Service) Gen() (bool, error) {
	file := service.GetTarget()

	content, err1 := service.BuildContent()

	if err1 != nil {
		return false, err1
	}

	return service.Write(file, content)
}
