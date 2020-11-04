package web

import "net/http"

/***
路由模块
*/

func Router() {
	http.HandleFunc("/Login", Login)
	//http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("static/pages"))))
	//http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("static/js"))))
	//web.Router()
	//http.ListenAndServe(":8080", nil)
}

/**
启动web 程序
*/
func StartWeb() {
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("static/pages"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("static/js"))))
	Router()
	http.ListenAndServe(":8080", nil)
}
