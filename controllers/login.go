package controllers

import (
	"github.com/sycdtk/bobi/logger"
)

//用户登录
func Login(w http.ResponseWriter, r *http.Request) {
	session := sessionManagerStart(w, r)
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	if len(username) > 0 && len(password) > 0 && session != nil {
		logger.Info(username, password, session.ID())
	}
}
