package internal

import (
	"gin-code-generator/internal/pkg/util"
	"time"
)

type Router struct {
	Base
}

type RouterTemplateData struct {
	TemplateData
	NameSpace string
	Version   string
	LowerName string
	WithCurd  bool
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

func (router *Router) BuildContent() (string, error) {
	name := router.BuildName()
	namespace, err := util.GetNameSpace(router.Config.Path + "/")
	if err != nil {
		return "", err
	}
	return util.ParseTemplateFromAssets(router.TemplateFile, &RouterTemplateData{
		TemplateData: TemplateData{
			Author: router.Config.Author,
			Date:   router.Date,
			Name:   name,
		},
		WithCurd:  router.Config.WithCurd,
		NameSpace: namespace,
		Version:   router.Config.ApiVersion,
		LowerName: util.LowerFirst(name),
	})
}

func (router *Router) Gen() (bool, error) {
	file := router.GetTarget()

	content, err := router.BuildContent()

	if err != nil {
		return false, err
	}

	return router.Write(file, content)
}
