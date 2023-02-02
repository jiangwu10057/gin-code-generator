package internal

import (
	"archive/zip"
	"gin-code-generator/assets"
	"gin-code-generator/internal/pkg/util"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

type Project struct {
	Base
	Namespace string
}

func NewProject(config Config) *Project {
	return &Project{
		Base: Base{
			Date:         time.Now().Format("2006/01/02 15:04"),
			TargetPath:   config.Path,
			TemplateFile: "assets/project.zip",
			FileSuffix:   "",
			Config:       config,
		},
		Namespace: config.Name,
	}
}

func (project *Project) Gen() (bool, error) {

	zip, err := project.generateTempZip()

	if err != nil {
		return false, err
	}

	err = project.unzip(zip, project.TargetPath)
	if err != nil {
		return false, err
	}

	err = project.removeZip(zip)
	if err != nil {
		return false, err
	}

	return true, err
}

func (project *Project) generateTempZip() (string, error) {
	data, err := assets.Asset(project.TemplateFile)

	if err != nil {
		return "", err
	}

	zip := "./template.zip"
	err = project.writeZip(zip, data)
	if err != nil {
		return "", err
	}

	return zip, nil
}

func (project *Project) writeZip(file string, data []byte) error {
	return ioutil.WriteFile(file, data, 0666)
}

func (project *Project) removeZip(file string) error {
	return os.Remove(file)
}

func (project *Project) unzip(archive, target string) error {
	reader, err := zip.OpenReader(archive)
	if err != nil {
		return err
	}

	defer reader.Close()

	if err := os.MkdirAll(target, 0755); err != nil {
		return err
	}

	for _, file := range reader.File {
		path := filepath.Join(target, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(path, file.Mode())
			continue
		}

		dir := filepath.Dir(path)
		if len(dir) > 0 {
			if _, err = os.Stat(dir); os.IsNotExist(err) {
				err = os.MkdirAll(dir, 0755)
				if err != nil {
					return err
				}
			}
		}

		fileReader, err := file.Open()
		if err != nil {
			return err
		}
		defer fileReader.Close()

		body, err := ioutil.ReadAll(fileReader)

		if err != nil {
			return err
		}

		content := string(body[:])

		content = util.ReplaceAll(content, "NAMESPACE", project.Namespace)

		err = ioutil.WriteFile(path, []byte(content), 0644)
		if err != nil {
			return err
		}
	}

	return nil
}
