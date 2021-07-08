package main

import (
	// "flag"
	"flag"
	"fmt"
	"log"

	// "log"
	"os"

	"gin-code-generator/internal"
	pkgconfig "gin-code-generator/internal/pkg/config"
	"gin-code-generator/internal/pkg/model"
)

var (
	VERSION  = "1.0.1"
	exitCode int
)

// Possible exit status codes.
const (
	ExitOk         = iota // nothing has changed or needs to change
	ExitDiff              // something has changed or needs to change
	ExitWrongUsage        // error in usage
)

func start(config internal.Config) {
	if config.WithCurd {
		_init()
	}
	switch config.Module {
	case "project":
		{
			generate(config)
		}
	case "model", "service", "router", "api":
		{
			generate(config)
		}
	case "struct":
		{
			_init()
			generate(config)
		}
	case "quick":
		{
			for _, value := range []string{"model", "service", "router", "api"} {
				config.Module = value
				generate(config)
			}
		}
	default:
		{
			fmt.Println("module is unexpected,except values:'project','api','model','service','router','api','quick','struct'")
			os.Exit(ExitWrongUsage)
		}
	}
}

func doGenerate(generator internal.Generator) {
	_, err := generator.Init()

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(exitCode)
	}

	_, err = generator.Gen()

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(exitCode)
	}
}

func generate(config internal.Config) {
	var generator internal.Generator

	switch config.Module {
	case "project":
		{
			generator = internal.NewProject(config)
		}
	case "model":
		{
			generator = internal.NewModel(config)

			if config.WithCurd {
				doGenerate(internal.NewRequestModel(config))
				doGenerate(internal.NewResponseModel(config))
			}
		}
	case "service":
		{
			generator = internal.NewService(config)
		}
	case "router":
		{
			generator = internal.NewRouter(config)
		}
	case "api":
		{
			generator = internal.NewApi(config)
		}
	case "struct":
		{
			generator = internal.NewTableStruct(config)
		}
	}

	doGenerate(generator)

	if config.WithTest {
		module := config.Module
		config.Module = "test"
		generator = internal.NewTest(config, module)
		doGenerate(generator)
	}
}

func _init() {
	fullConfig, err := pkgconfig.LoadConfig()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(ExitOk)
	}
	err = model.InitOrm(fullConfig)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(ExitOk)
	}
}

func main() {
	log.SetFlags(0)
	flag.Usage = func() {
		fmt.Println("Gin Code Generator [options]")
		fmt.Println()
		flag.PrintDefaults()
	}

	config := internal.Config{}

	pwd, _ := os.Getwd()
	author, _ := os.Hostname()

	flag.StringVar(&config.Module, "module", "project", "which module you want to generate.\noption:'project','model','service','router','api','quick','struct'")
	flag.StringVar(&config.Path, "path", pwd, "project root path,default is current path")
	flag.StringVar(&config.Author, "author", author, "code author,default is computer name")
	flag.StringVar(&config.Name, "name", "", "module name or project name")
	flag.StringVar(&config.ApiVersion, "apiv", "", "api version")
	flag.StringVar(&config.Tags, "tags", "", "\"tags\" is api swagger notes")
	flag.BoolVar(&config.Force, "force", false, "use the option if file exist,it will replace. and if the path don't exist,it will create.")
	flag.BoolVar(&config.WithTest, "withtest", false, "do you want to generate test file in the same time ? default is no")
	flag.BoolVar(&config.WithCurd, "withcurd", false, "do you want to generate CURD API in the same time ? default is no")
	showVersion := flag.Bool("v", false, "print version")

	flag.Parse()

	if *showVersion {
		fmt.Println(VERSION)
		os.Exit(ExitOk)
	}
	if flag.NFlag() == 0 {
		log.Println("too few arguments")
		flag.Usage()
		os.Exit(ExitWrongUsage)
	}

	start(config)

	fmt.Println("ok")

	os.Exit(exitCode)
}
