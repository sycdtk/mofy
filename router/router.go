package router

import (
	"net/http"

	"github.com/sycdtk/bobi/web/restful"
	"github.com/sycdtk/mofy/controllers"
)

func init() {
	//注册静态文件服务
	restful.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./xin"))), http.MethodGet)
	// http.Handle("/image/", http.StripPrefix("/image/", http.FileServer(http.Dir("./static/image"))))
	// http.Handle("/dist/", http.StripPrefix("/dist/", http.FileServer(http.Dir("./static/dist"))))
	// http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./static/css"))))
	// http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./static/js"))))
	// http.Handle("/vendor/", http.StripPrefix("/vendor/", http.FileServer(http.Dir("./static/vendor"))))

	// //注册controllers
	restful.HandleFunc("/login", controllers.Login, http.MethodGet, false)
	// http.HandleFunc("/logout", controllers.Logout)
	// http.HandleFunc("/index", controllers.Index)
	// http.HandleFunc("/notFound", controllers.NotFound)
	// http.HandleFunc("/error", controllers.ErrorPage)

	// //用户配置
	// http.HandleFunc("/userManager", controllers.UserManager)
	// http.HandleFunc("/userList", controllers.UserList)
	// http.HandleFunc("/addUser", controllers.AddUser)
	// http.HandleFunc("/delUser", controllers.DelUser)

	// http.HandleFunc("/", controllers.Default)
}
