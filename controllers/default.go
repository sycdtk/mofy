package controllers

import (
	"html/template"
	"net/http"

	"github.com/sycdtk/bobi/web"
	"github.com/sycdtk/bobi/web/session"
)

//调用session manager start
func sessionManagerStart(w http.ResponseWriter, r *http.Request) (session session.Session) {
	return web.SessionManager.Start(w, r)
}

//调用session manager destroy
func sessionManagerDestroy(w http.ResponseWriter, r *http.Request) {
	web.SessionManager.Destroy(w, r)
}

//默认页面
func Default(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("views/default.tpl"))
	t.Execute(w, nil)
}

//404页面
func NotFound(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("views/default.tpl"))
	t.Execute(w, nil)
}

//错误页面
func ErrorPage(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("views/default.tpl"))
	t.Execute(w, nil)
}
