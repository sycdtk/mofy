package controllers

import (
	"fmt"
	"net/http"

	// "net/url"
	"strings"

	"github.com/sycdtk/bobi/logger"
)

//用户登录
func Login(w http.ResponseWriter, r *http.Request, paramsMap map[string]string) interface{} {

	var username, password string
	logger.Info(r.URL)
	// if queryForm, err := url.ParseRequestURI(r.URL.RawQuery); err == nil {
	// 	username = queryForm["username"].
	// 	password = queryForm["password"]
	// }

	logger.Info(username, password)

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	if len(username) > 0 && len(password) > 0 {
		logger.Info(username, password)
		return []string{"aaa", "bbb"}
	}
	return nil
}
