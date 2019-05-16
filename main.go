package main

import (
	"github.com/sycdtk/bobi/config"
	"github.com/sycdtk/bobi/web/restful"

	_ "github.com/sycdtk/mofy/router" //初始化路由
)

func main() {
	restful.ListenAndServe(config.Read("web", "port")) //启动监听服务
}
