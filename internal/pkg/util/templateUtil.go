package util

import (
	"bytes"
	"text/template"
)

func ParseTemplate(templateFilePath string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templateFilePath)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	err = t.Execute(buf, data)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
