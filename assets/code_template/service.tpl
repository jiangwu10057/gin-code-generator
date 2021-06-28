/**
* @Description {{.Name}}服务类
* @Author {{.Author}}
* @Date  {{.Date}}
**/
package service

import (
    {{if .WithCurd}}
    "{{.NameSpace}}/model/request"
    "{{.NameSpace}}/model/response"
    {{end}}
)

type {{.Name}}Service struct {

}

{{if .WithCurd}}
// 获取列表
func (_ *{{.Name}}Service) GetList(info request.Get{{.Name}}Req) (list []response.Get{{.Name}}Resp, total int64, err error) {
	return list, total, err
}

// 删除一条记录  逻辑删除
func (_ *{{.Name}}Service) Delete{{.Name}}(id string) error {
	return nil
}

//新增一条记录
func (_ *{{.Name}}Service) Create{{.Name}}(info request.Add{{.Name}})  (request.Add{{.Name}}, error) {
    return info,nil
}

//更新一条记录
func (_ *{{.Name}}Service) Update{{.Name}}(r request.Update{{.Name}}) error {
	return nil
}

{{end}}