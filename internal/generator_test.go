package internal

import (
	"os"
	"time"
	"gin-code-generator/internal"
	"github.com/stretchr/testify/assert"
	"testing"

)

func TestNewProject(t *testing.T) {
	config := internal.Config{
		Module: "",
		Name:   "",
		Path:   "",
		Author: "",
	}
	
	project := internal.NewProject(config)
	assert.Equal(t, "/", project.TargetPath)
	assert.Equal(t, "/assets/", project.TemplateFile)
}

func TestNewService(t *testing.T) {
	config := internal.Config{
		Module: "",
		Name:   "",
		Path:   "",
		Author: "",
	}
	
	service := internal.NewService(config)
	assert.Equal(t, "/service/", service.TargetPath)
	assert.Equal(t, "/assets/code_template/service.tpl", service.TemplateFile)
}

// func TestNewTemplateData(t *testing.T) {
// 	config := internal.Config{
// 		Module: "",
// 		Name:   "",
// 		Path:   "",
// 		Author: "气昂昂",
// 	}
	
// 	service := internal.NewService(config)
// 	templateData := internal.NewTemplateData(service)
// 	assert.Equal(t, config.Author, templateData.Author)
// }

func TestNewRouter(t *testing.T) {
	config := internal.Config{
		Module: "",
		Name:   "",
		Path:   "",
		Author: "",
	}
	
	router := internal.NewRouter(config)
	assert.Equal(t, "/router/", router.TargetPath)
	assert.Equal(t, "/assets/code_template/router.tpl", router.TemplateFile)
}

func TestNewModel(t *testing.T) {
	config := internal.Config{
		Module: "",
		Name:   "",
		Path:   "",
		Author: "",
	}
	
	model := internal.NewModel(config)
	assert.Equal(t, "/model/", model.TargetPath)
	assert.Equal(t, "/assets/code_template/model.tpl", model.TemplateFile)
}

func TestNewTest(t *testing.T) {
	config := internal.Config{
		Module: "",
		Name:   "",
		Path:   "",
		Author: "",
	}
	
	test := internal.NewTest(config)
	assert.Equal(t, "/", test.TargetPath)
	assert.Equal(t, "/assets/code_template/test.tpl", test.TemplateFile)
}

func TestCheckModuleFail(t *testing.T) {
	Config := internal.Config{
		Module: "",
		Name:   "",
		Path:   "",
		Author: "",
	}

	base := internal.Base{
		Config: Config,
	}

	status, err := base.CheckModule()
	assert.Equal(t, false, status)
	assert.Equal(t, "module cannot be empty", err.Error())

	Config = internal.Config{
		Module: "module",
		Name:   "",
		Author: "",
		Path:   "",
	}

	base = internal.Base{
		Config: Config,
	}

	status, err = base.CheckModule()
	assert.Equal(t, false, status)
	assert.Equal(t, "module is unexpected,except values:'project','api','model','service','router','test'", err.Error())
}

func TestCheckModuleSuccsess(t *testing.T) {
	Config := internal.Config{
		Module: "api",
		Name:   "",
		Path:   "",
		Author: "",
	}

	base := internal.Base{
		Config: Config,
	}

	status, err := base.CheckModule()
	assert.Equal(t, true, status)
	assert.Equal(t, nil, err)
}

func TestInit(t *testing.T) {
	Config := internal.Config{
		Module: "api",
		Name:   "",
		Path:   "",
		Author: "",
	}

	base := internal.Base{
		Config: Config,
	}

	status, err := base.Init()
	assert.Equal(t, false, status)
	assert.Equal(t, "name cannot be empty", err.Error())

	Config = internal.Config{
		Module: "api",
		Name:   "account",
		Path:   "",
		Author: "",
	}

	base = internal.Base{
		Config: Config,
	}

	status, err = base.Init()
	assert.Equal(t, true, status)
	assert.Equal(t, nil, err)
}

func TestBuildName(t *testing.T) {
	Config := internal.Config{
		Module: "api",
		Name:   "account",
		Path:   "",
		Author: "",
	}

	base := internal.Base{
		Config: Config,
		Date:         "2021/06/18 11:10",
		TargetPath:   "/router/",
		TemplateFile: "/assets/code_template/router.tpl",
		FileSuffix:   "Router.go",
	}

	name := base.BuildName()

	assert.Equal(t, "Account", name)
}

// func TestBuildContent(t *testing.T) {
// 	Config := internal.Config{
// 		Module: "api",
// 		Name:   "account",
// 		Path:   "",
// 		Author: "",
// 	}

// 	base := internal.Base{
// 		Config: Config,
// 		Date:         "2021/06/18 11:10",
// 		TargetPath:   "/router/",
// 		TemplateFile: "/assets/code_template/router.tpl",
// 		FileSuffix:   "Router.go",
// 	}

// 	content,err := base.BuildContent()

// 	assert.Equal(t, "", content)
// 	assert.Equal(t, nil, err)
// }

func TestServiceBuildContent(t *testing.T) {
	Config := internal.Config{
		Module: "service",
		Name:   "account",
		Path:   "/golang/",
		Author: "test",
	}
	var generator internal.Generator
	generator = internal.NewService(Config)
	
	content, err := generator.BuildContent()
	date := time.Now().Format("2006/01/02 15:04")
	assert.Equal(t, "/**\r\n* @Author test\r\n* @Date  "+date+"\r\n**/\r\npackage service\r\n\r\nimport (\r\n\r\n)\r\n\r\ntype AccountService struct {\r\n}", content)
	assert.NoError(t, err)
}

func TestModelBuildContent(t *testing.T) {
	Config := internal.Config{
		Module: "model",
		Name:   "account",
		Path:   "/golang/",
		Author: "test",
	}
	var generator internal.Generator
	generator = internal.NewModel(Config)
	
	content, err := generator.BuildContent()
	date := time.Now().Format("2006/01/02 15:04")
	assert.Equal(t, "/**\r\n* @Author test\r\n* @Date  "+date+"\r\n**/\r\npackage model\r\n\r\ntype AccountModel struct {\r\n\r\n}\r\n\r\nfunc (AccountModel) TableName() string {\r\n\treturn \"ACCOUNT\"\r\n}", content)
	assert.NoError(t, err)
}

func TestRouterBuildContent(t *testing.T) {
	Config := internal.Config{
		Module: "router",
		Name:   "account",
		Path:   "/golang/",
		Author: "test",
	}
	var generator internal.Generator
	generator = internal.NewRouter(Config)
	
	content, err := generator.BuildContent()
	date := time.Now().Format("2006/01/02 15:04")
	assert.Equal(t, "/**\r\n* @Author test\r\n* @Date  "+date+"\r\n**/\r\n\r\npackage router\r\n\r\nimport (\r\n\t\"github.com/gin-gonic/gin\"\r\n)\r\n\r\nfunc AccountRouter(r *gin.RouterGroup) {\r\n\t\r\n}\r\n", content)
	assert.NoError(t, err)
}

func TestTestBuildContent(t *testing.T) {
	Config := internal.Config{
		Module: "test",
		Name:   "account",
		Path:   "/golang/",
		Author: "test",
	}
	var generator internal.Generator
	generator = internal.NewTest(Config)
	
	content, err := generator.BuildContent()
	date := time.Now().Format("2006/01/02 15:04")
	assert.Equal(t, "/**\r\n* @Author test\r\n* @Date  "+date+"\r\n**/\r\n\r\npackage \r\n\r\nimport (\r\n\t\"testing\"\r\n\t\"github.com/stretchr/testify/assert\"\r\n)\r\n\r\nfunc Test[](t *testing.T){\r\n\r\n}", content)
	assert.NoError(t, err)
}

func TestGetTarget(t *testing.T) {
	Config := internal.Config{
		Module: "router",
		Name:   "account",
		Path:   "/golang",
		Author: "test",
	}
	var generator internal.Generator
	generator = internal.NewRouter(Config)
	target := generator.GetTarget()
	assert.Equal(t, "/golang/router/AccountRouter.go", target)

}

func TestWrite(t *testing.T) {
	Config := internal.Config{
		Module: "router",
		Name:   "account",
		Path:   "/golang",
		Author: "test",
	}
	var generator internal.Generator
	generator = internal.NewRouter(Config)
	target := generator.GetTarget()
	os.Remove(target)
	content, _ := generator.BuildContent()
	result,err := generator.Write(target, content)
	assert.Equal(t, true, result)
	assert.NoError(t, err)
}

func TestGen(t *testing.T) {
	Config := internal.Config{
		Module: "service",
		Name:   "account",
		Path:   "/golang",
		Author: "test",
	}
	var generator internal.Generator
	generator = internal.NewService(Config)
	target := generator.GetTarget()
	os.Remove(target)
	result,err := generator.Gen()
	assert.Equal(t, true, result)
	assert.NoError(t, err)
}
