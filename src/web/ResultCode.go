package web

/**
返回给前端的json字符串基本格式
*/
type ResultJson struct {
	Code  uint8
	Value string
	Msg   string
}
