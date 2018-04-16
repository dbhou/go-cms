package main

import (
	_ "cms/routers"
	//	"mime"
	"os"

	. "cms/lib"
	"cms/models"

	"github.com/astaxie/beego"
)

func main() {
	//初始化
	initialize()

	beego.Run()
}

func initialize() {
	//	mime.AddExtensionType(".css", "text/css")
	//判断初始化参数
	//	initArgs()

	models.Connect()

	beego.AddFuncMap("stringsToJson", StringsToJson)
}

func initArgs() {
	args := os.Args
	for _, v := range args {
		if v == "-syncdb" {
			models.Syncdb()
			os.Exit(0)
		}
	}
}
