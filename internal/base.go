package internal

import (
	"errors"
	"fmt"
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
	Force      bool
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

func (base *Base) forceCreate(path string, file string) error {

	if !util.CheckFileIsExist(path) {
		err := os.MkdirAll(path, 0666)
		if err != nil {
			return err
		}
	}
	if util.CheckFileIsExist(path + file) {
		err := os.Remove(path + file)
		if err != nil {
			return err
		}
	}

	return nil
}

func (base *Base) GetTarget() string {
	path := base.Config.Path + base.TargetPath
	file := base.BuildName() + base.FileSuffix

	if base.Config.Force {
		err := base.forceCreate(path, file)
		if err != nil {
			fmt.Print(err.Error())
		}
	}

	return path + file
}

func (base *Base) Write(file string, content string) (bool, error) {
	if util.CheckFileIsExist(file) {
		return true, errors.New(file + " already exists")
	}

	f, err := os.Create(file)

	if err != nil {
		return false, err
	}

	_, err = io.WriteString(f, content)

	if err == nil {
		util.GoFileFormat(file)
	}

	return err == nil, err
}

func (base *Base) Gen() (bool, error) {
	file := base.GetTarget()

	content, err := base.BuildContent()

	if err != nil {
		return false, err
	}

	return base.Write(file, content)
}
