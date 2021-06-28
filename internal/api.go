package internal

import (
	"errors"
	"gin-code-generator/internal/pkg/util"
	"time"
)

type Api struct {
	Base
}

type ApiTemplateData struct {
	TemplateData
	Version   string
	NameSpace string
	LowerName string
	Tags      string
	WithCurd  bool
}

func NewApi(config Config) *Api {
	target := "/api/"
	if len(config.ApiVersion) > 0 {
		target = "/api/" + config.ApiVersion + "/"
	}
	return &Api{
		Base{
			Date:         time.Now().Format("2006/01/02 15:04"),
			TargetPath:   target,
			TemplateFile: "assets/code_template/api.tpl",
			FileSuffix:   "Api.go",
			Config:       config,
		},
	}
}

func (api *Api) Init() (bool, error) {
	if api.Config.Name == "" {
		return false, errors.New("name cannot be empty")
	}

	if api.Config.Tags == "" {
		return false, errors.New("tags cannot be empty")
	}

	return true, nil
}

func (api *Api) BuildContent() (string, error) {
	namespace, err := util.GetNameSpace(api.Config.Path + "/")
	if err != nil {
		return "", err
	}

	name := api.BuildName()

	data := &ApiTemplateData{
		TemplateData: TemplateData{
			Author: api.Config.Author,
			Date:   api.Date,
			Name:   name,
		},
		WithCurd:  api.Config.WithCurd,
		NameSpace: namespace,
		Version:   api.Config.ApiVersion,
		LowerName: util.LowerFirst(name),
		Tags:      api.Config.Tags,
	}

	return util.ParseTemplateFromAssets(api.TemplateFile, data)
}

func (api *Api) Gen() (bool, error) {
	file := api.GetTarget()

	content, err := api.BuildContent()

	if err != nil {
		return false, err
	}

	return api.Write(file, content)
}
