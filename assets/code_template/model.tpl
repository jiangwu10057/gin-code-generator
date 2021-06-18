/**
* @Author {{.Author}}
* @Date  {{.Date}}
**/
package model

type {{.Name}}Model struct {

}

func ({{.Name}}Model) TableName() string {
	return "{{.TableName}}"
}