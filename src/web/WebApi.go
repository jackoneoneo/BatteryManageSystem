package web

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/**
该模块实现登录界面的业务
*/

// 登录接口
func Login(w http.ResponseWriter, r *http.Request) {
	resultJson := ResultJson{Code: 2, Value: "222", Msg: "3333"}
	b, err := json.Marshal(resultJson)
	if err != nil {
		fmt.Println("Umarshal failed:", err)
	}
	fmt.Fprintf(w, string(b))
}
