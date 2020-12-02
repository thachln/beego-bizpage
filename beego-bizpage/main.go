package main

import (
	_ "beego-bizpage/routers"
	beego "github.com/astaxie/beego/server/web"
)

func main() {
	beego.Run()
}

