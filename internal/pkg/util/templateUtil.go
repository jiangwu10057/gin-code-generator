package util

import (
	"bytes"
	"gin-code-generator/assets"
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

func ParseTemplateFromAssets(templateFilePath string, data interface{}) (string, error) {

	b, err := assets.Asset(templateFilePath) // 根据地址获取对应内容
	if err != nil {
		return "", err
	}

	t, err1 := template.New("tmp").Parse(string(b))
	if err1 != nil {
		return "", err1
	}
	buf := new(bytes.Buffer)
	err = t.Execute(buf, data)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
