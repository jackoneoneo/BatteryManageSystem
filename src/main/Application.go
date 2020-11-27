package main

import (
	"communication"
	"web"
)

/*

 */

type Name struct {
	dd string
}

/**
初始化日志文件
*/
func init() {
	//println("系统初始化入口")
	////fmt.Println("main init")
	//file := "./" + "bms.log"
	//logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	//if err != nil {
	//	panic(err)
	//}
	//log.SetOutput(logFile) // 将文件设置为log输出的文件
	//log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
}

type Node struct {
	priority int
	value    string
}

func (this *Node) Less(other interface{}) bool {
	return this.priority < other.(*Node).priority
}

//执行命令
// 1. cd ./src/main  2. D:\goRepository\src\github.com\mitchellh\gox\gox.exe
// 3. D:\goRepository\src\github.com\mitchellh\gox\gox.exe -os "linux" 只编译linux版本
// 4.D:\goRepository\src\github.com\mitchellh\gox\gox.exe -osarch="linux/arm" 只编译arm版本
// 程序的入口
//func main() {

// q := priority_queue.New()
//q.Push(&Node{priority: 8, value: "8"})
//q.Push(&Node{priority: 7, value: "7"})
//q.Push(&Node{priority: 9, value: "9"})
//q.Push(&Node{priority: 2, value: "2"})
//q.Push(&Node{priority: 4, value: "4"})
//q.Push(&Node{priority: 3, value: "3"})
//q.Push(&Node{priority: 5, value: "5"})
//x := q.Top().(*Node)
//fmt.Println(x.priority, x.value)
//
//for q.Len() > 0 {
//	x = q.Pop().(*Node)
//	fmt.Println(x.priority, x.value)
//}
//fmt.Print("hello world")
//value := 0x86
//fmt.Print(value)
//http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("static/pages"))))
//http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("static/js"))))
//web.Router()
//http.ListenAndServe(":8080", nil)
//
//resultJson := util.ResultJson{Code: 2, Value: "222", Msg: "3333"}
//b, err := json.Marshal(resultJson)
//if err != nil {
//	fmt.Println("Umarshal failed:", err)
//}

//}

func main() {
	go communication.CBmsServer()
	web.StartWeb()
}
