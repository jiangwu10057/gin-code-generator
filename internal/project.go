package internal

import (
	"time"
)

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
