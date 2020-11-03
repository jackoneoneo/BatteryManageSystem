package web

import "net/http"

/***
路由模块
*/

func Router() {
	http.HandleFunc("/Login", Login)
}
