package main

import (
	// "flag"
	"fmt"
	// "log"
	"os"

	"gin-code-generator/internal"
	pkgconfig "gin-code-generator/internal/pkg/config"
	"gin-code-generator/internal/pkg/model"
)

var (
	VERSION  = "0.0.2"
	exitCode int
)

// Possible exit status codes.
const (
	ExitOk         = iota // nothing has changed or needs to change
	ExitDiff              // something has changed or needs to change
	ExitWrongUsage        // error in usage
)


func start(config internal.Config) {
	switch config.Module {
		case "project":{
			fmt.Println("to be continue")
			os.Exit(ExitWrongUsage)
		}
		case "model":{
			generate(config)
		}
		case "service":{
			generate(config)
		}
		case "router":{
			generate(config)
		}
	    case "quick":{
			for _,value := range []string{"model", "service", "router"}{
				config.Module = value
				generate(config)
			}
		}
		default:{
			fmt.Println("module is unexpected,except values:'project','api','model','service','router','quick'")
			os.Exit(ExitWrongUsage)
		}
	}
}

func generate(config internal.Config){
	var generator internal.Generator

	switch config.Module {
		case "model":{
			generator = internal.NewModel(config)
		}
		case "service":{
			generator = internal.NewService(config)
		}
		case "router":{
			generator = internal.NewRouter(config)
		}
	}

	_, err := generator.Init()

	if err != nil{
		fmt.Println(err.Error())
		os.Exit(exitCode)
	}

	_,err = generator.Gen()

	if err != nil{
		fmt.Println(err.Error())
		os.Exit(exitCode)
	}

	if config.WithTest {
		module := config.Module
		config.Module = "test"
		generator = internal.NewTest(config, module)
		_, err = generator.Init()

		if err != nil{
			fmt.Println(err.Error())
			os.Exit(exitCode)
		}
	
		_,err = generator.Gen()
	
		if err != nil{
			fmt.Println(err.Error())
			os.Exit(exitCode)
		}
	}
}

func init() {
	fullConfig, err := pkgconfig.LoadConfig()
	if err != nil{
		fmt.Println(err.Error())
		os.Exit(ExitOk)
	}
	err = model.InitOrm(fullConfig)

	if err != nil{
		fmt.Println(err.Error())
		os.Exit(ExitOk)
	}
}

func main() {
	tableName := "user_type_config"
	operator := model.NewColumnOperator()
	var getter model.ColumnGetter
	getter = model.NewOracleColumnGetter()
	err := operator.SetStrategy(getter).Get(tableName)
	if err != nil{
		fmt.Println(err.Error())
		os.Exit(ExitOk)
	}

	builder := model.NewOracleModelBuilder(tableName, getter)
	str, err1 := builder.Create()
	if err1 != nil{
		fmt.Println(err1.Error())
		os.Exit(ExitOk)
	}

	fmt.Println(str)
	

	// fmt.Println(columns)
	// log.SetFlags(0)
	// flag.Usage = func() {
	// 	fmt.Println("Gin Code Generator [options]")
	// 	fmt.Println()
	// 	flag.PrintDefaults()
	// }

	// config := internal.Config{}

	// pwd, _ := os.Getwd()
	// author, _ := os.Hostname()

	// flag.StringVar(&config.Module, "module", "project", "which module you want to generate.\noption:'project','model','service','router','quick'")
	// flag.StringVar(&config.Path, "path", pwd, "project root path,default is current path")
	// flag.StringVar(&config.Author, "author", author, "code author,default is computer name")
	// flag.StringVar(&config.Name, "name", "", "name for module")
	// flag.BoolVar(&config.WithTest, "withtest", false, "do you want to generate test file in the same time?defualt is no")
	// showVersion := flag.Bool("v", false, "print version")

	// flag.Parse()

	// if *showVersion {
	// 	fmt.Println(VERSION)
	// 	os.Exit(ExitOk)
	// }
	// if flag.NFlag() == 0 {
	// 	log.Println("too few arguments")
	// 	flag.Usage()
	// 	os.Exit(ExitWrongUsage)
	// }
	
	// start(config)
	
	// fmt.Println("ok")

	
	// os.Exit(exitCode)
}
