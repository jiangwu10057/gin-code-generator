/**
* @Description {{.Name}}表API接口
* @Author {{.Author}}
* @Date  {{.Date}}
**/

package {{if gt (len .Version) 0}}{{.Version}}{{else}}api{{end}}

import (
	{{if .WithCurd}}
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	{{end}}
	{{if .WithCurd}}
	"{{.NameSpace}}/global"
	"{{.NameSpace}}/global/response"
	"{{.NameSpace}}/model"
	"{{.NameSpace}}/model/request"
	resp "{{.NameSpace}}/model/response"
	{{end}}
	"{{.NameSpace}}/service"
	{{if .WithCurd}}
	"{{.NameSpace}}/utils"
	{{end}}
)

var {{.LowerName}}Service service.{{.Name}}Service

{{if .WithCurd}}
// @Tags {{.Tags}}
// @Summary 获取
// @Security ApiKeyAuth
// @Description 
// @Produce application/json
// @Param data query request.Get{{.Name}}Req true ""
// @Success 200 {string} string "{"code":"0","data":{},"message":"获取成功"}"
// @Failure 400 {string} string "{"code":"-1","data":{},"message":"获取失败"}"
// @Router /api/{{if gt (len .Version) 0}}{{.Version}}/{{end}}{{.LowerName}}/get{{.Name}} [get]
func Get{{.Name}}(c *gin.Context) {
	var r request.Get{{.Name}}Req

	var err error
	r.Page, err = strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	r.PageSize, err = strconv.Atoi(c.DefaultQuery("pageSize", "5"))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := {{.LowerName}}Service.GetList(r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(resp.PageResult{
			List:     list,
			Total:    total,
			Page:     r.Page,
			PageSize: r.PageSize,
		}, c)
	}
}

// @Tags {{.Tags}}
// @Summary 删除
// @Security ApiKeyAuth
// @Description
// @Produce application/json
// @Param data body request.Delete{{.Name}}Req true ""
// @Success 200 {string} string "{"code":"0","data":"success","message":"调用成功"}"
// @Failure 500 {string} string "{"code":"-1","data":{},"message":"失败的说明"}"
// @Router /api/{{if gt (len .Version) 0}}{{.Version}}/{{end}}{{.LowerName}}/delete{{.Name}} [delete]
func Delete{{.Name}}(c *gin.Context) {
	var r request.Delete{{.Name}}Req
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = {{.LowerName}}Service.Delete{{.Name}}(r.ID)//请修改字段名
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	addOpsOperationLog(c, "", "", 2, 5) //请修改操作日志文案
	response.OkWithMessage("", c) //请修改返回文案
}

func addOpsOperationLog(c *gin.Context, detail string, warnInfo string, moduleType int, operateType int) {
	claim := global.GetMemberBasicData(c)
	OperateLog := model.OpsOperateLogModel{
		LogId:         utils.GetUUID(),
		OperateDetail: detail,
		ProductId:     claim.ProductIds,
		ClientIp:      strings.Split(c.Request.RemoteAddr, ":")[0],
		UserId:        claim.UserId,
		UserName:      claim.UserName,
		ModuleType:    moduleType,
		OperateType:   operateType,
		WarnInfo:      warnInfo,
	}
	//操作日志
	err := OpsOperateLogService.OperateLog(OperateLog)
	if err != nil {
		global.LOGGER.Error(fmt.Sprintf("操作日志写入失败，%v", err))
	}
}

// @Tags {{.Tags}}
// @Summary 新增
// @Security ApiKeyAuth
// @Produce application/json
// @Param data body request.Add{{.Name}}Req true ""
// @Success 200 {string} string "{"code":"0","data":{},"message":"添加成功"}"
// @Failure 500 {string} string "{"code":"-1","data":{},"message":"失败的说明"}"
// @Router /api/{{if gt (len .Version) 0}}{{.Version}}/{{end}}{{.LowerName}}/add{{.Name}} [post]
func Add{{.Name}}(c *gin.Context) {
	var r request.Add{{.Name}}Req
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	retModel, err := {{.LowerName}}Service.Create{{.Name}}(r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		addOpsOperationLog(c, "", "", 2, 3)//请修改操作日志文案
		response.OkDetailed(retModel, "添加成功", c)
	}
}

// @Tags {{.Tags}}
// @Summary 修改
// @Description
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.Update{{.Name}}Req true ""
// @Success 200 {string} string "{"code":"0","data":{},"message":"修改成功"}"
// @Failure 400 {string} string "{"code":"-1","data":{},"message":"失败的说明"}"
// @Router /api/{{if gt (len .Version) 0}}{{.Version}}/{{end}}{{.LowerName}}/update{{.Name}} [put]
func Update{{.Name}}(c *gin.Context) {
	var r request.Update{{.Name}}Req
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = {{.LowerName}}Service.Update{{.Name}}(r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		addOpsOperationLog(c, "", "", 2, 4)//请修改操作日志文案
		response.OkWithMessage("修改成功", c)
	}
}
{{end}}