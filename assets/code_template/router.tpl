/**
* @Description  {{.Name}}表API路由
* @Author {{.Author}}
* @Date  {{.Date}}
**/

package router

import (
	"github.com/gin-gonic/gin"

	{{if .WithCurd}}
	api "{{.NameSpace}}/api{{if gt (len .Version) 0}}/{{.Version}}{{end}}"
	{{end}}
)

func {{.Name}}Router(r *gin.RouterGroup) {
	{{if .WithCurd}}
	ur := r.Group("{{if gt (len .Version) 0}}{{.Version}}/{{end}}{{.LowerName}}")
	{
		ur.GET("get{{.Name}}", api.Get{{.Name}})
		ur.POST("add{{.Name}}", api.Add{{.Name}})
		ur.PUT("update{{.Name}}", api.Update{{.Name}})
		ur.GET("delete{{.Name}}", api.Delete{{.Name}})
	}
	{{end}}
}
