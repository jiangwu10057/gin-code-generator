/**
* @Description {{.TableName}}表结构体
* @Author {{.Author}}
* @Date  {{.Date}}
**/
package model

type {{.Name}}Model struct {
{{.Fields}}
}

func ({{.Name}}Model) TableName() string {
	return "{{.TableName}}"
}