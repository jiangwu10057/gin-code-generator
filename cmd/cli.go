package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"gin-code-generator/internal"
)

var (
	VERSION  = "0.0.1"
	exitCode int
)

// Possible exit status codes.
const (
	ExitOk         = iota // nothing has changed or needs to change
	ExitDiff              // something has changed or needs to change
	ExitWrongUsage        // error in usage
)

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

	flag.StringVar(&config.Module, "module", "project", "which module you want to generate.\noption:'project','model','service','router','quick'")
	flag.StringVar(&config.Path, "path", pwd, "project root path,default is current path")
	flag.StringVar(&config.Author, "author", author, "code author,default is computer name")
	flag.StringVar(&config.Name, "name", "", "name for module")
	flag.BoolVar(&config.WithTest, "withtest", false, "do you want to generate test file in the same time?defualt is no")
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
	var generator internal.Generator 
	switch config.Module {
		case "project", "quick":{
			fmt.Println("to be continue")
			os.Exit(ExitWrongUsage)
		}
		case "model":{
			generator = internal.NewModel(config)
		}
		case "service":{
			generator = internal.NewService(config)
		}
		case "router":{
			generator = internal.NewRouter(config)
		}
		default:{
			fmt.Println("module is unexpected,except values:'project','api','model','service','router','quick'")
			os.Exit(ExitWrongUsage)
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
	fmt.Println("ok")
	os.Exit(exitCode)
}
