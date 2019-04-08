package main

import (
	"net/http"

	_ "github.com/sycdtk/mofy/router" //初始化路由
)

func main() {

	http.ListenAndServe(":8080", nil) //启动监听服务

}
