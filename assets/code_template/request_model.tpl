/**
* @Description {{.TableName}}表请求结构体
* @Author {{.Author}}
* @Date  {{.Date}}
**/

package request

type Get{{.Name}}Req struct {
	PageInfoModel
}

type Delete{{.Name}}Req struct {
{{.Fields}}
}

type Add{{.Name}}Req struct {
{{.Fields}}
}

type Update{{.Name}}Req struct {
{{.Fields}}
}