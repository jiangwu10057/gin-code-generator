package internal

import (
	"errors"
	"gin-code-generator/internal/pkg/util"
	"io"
	"os"
	"strings"
)

type Config struct {
	Module     string
	Name       string
	Path       string
	Author     string
	ApiVersion string
	Tags       string
	WithTest   bool
	WithCurd   bool
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

type TemplateData struct {
	Author string
	Date   string
	Name   string
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

func (base *Base) GetTarget() string {
	return base.Config.Path + base.TargetPath + base.BuildName() + base.FileSuffix
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

	content, err := base.BuildContent()

	if err != nil {
		return false, err
	}

	return base.Write(file, content)
}
