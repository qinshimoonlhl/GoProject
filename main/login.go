package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	//指定访问的路由
	http.HandleFunc("/login",login)
	//设定监听端口
	err := http.ListenAndServe(":9090",nil)
	if err != nil {
		log.Fatal("ListenAndServe: ",err)
	}
}

/**
  登录后端处理逻辑
*/
func login(rw http.ResponseWriter,request *http.Request){
	// 获取请求的方法 request.Method 字符串类型的变量，返回GET, POST,PUT等method信息。
	fmt.Printf("method: %v\n",request.Method)
	//根据 request.Method 来判断是显示登录界面还是处理登录逻辑
	if request.Method == "GET" {
		//使用template.ParseFiles 根据指定的文件创建模板示例
		t,_ := template.ParseFiles("web/login.html")
		//执行数据融合
		t.Execute(rw,nil)
	}
	if request.Method == "POST" {
		//请求的是登录数据，那么执行登录的逻辑判断
		//解析传过来的参数，默认不会解析，必须显示调用后服务器才会输出参数信息
		request.ParseForm()
		//这里的request.Form["username"]可以用request.FormValue("username")代替，那么就不需要显示调用  request.ParseForm
		fmt.Printf("username: %v\n",request.Form["username"])
		fmt.Printf("password: %v\n",request.Form["password"])
	}
}