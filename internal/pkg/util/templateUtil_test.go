package util

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"

	"gin-code-generator/internal/pkg/util"
)

type Data struct {
	Author    string
	Date      string
	Name      string
	Fields    string
	TableName string
}

func TestParseTemplate(t *testing.T) {
	currentPath, _ := os.Getwd()
	var templateFilePath string = currentPath + "/../../../assets/code_template/model.tpl"
	data := Data{
		Author:    "测试",
		Date:      "2021/06/17 16:05",
		Name:      "Account",
		Fields:    "AAAA",
		TableName: "ACCOUNT",
	}

	content, err := util.ParseTemplate(templateFilePath, data)
	assert.NoError(t, err)
	assert.Equal(t, "/**\r\n* @Author 测试\r\n* @Date  2021/06/17 16:05\r\n**/\r\npackage model\r\n\r\ntype AccountModel struct {\r\n\tAAAA\r\n}\r\n\r\nfunc (AccountModel) TableName() string {\r\n\treturn \"ACCOUNT\"\r\n}", content)
}
