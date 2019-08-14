package main

import (
	"net/http"
	"../route"

)


// 初始化参数
func init() {
	//初始化日志输出

}

func main() {
	//启动服务
	router := route.NewRouter()
	http.ListenAndServe(":9092", router)
}