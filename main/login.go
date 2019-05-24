package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {
	//指定访问的路由
	http.HandleFunc("/login", login)
	//设定监听端口
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

/**
  登录后端处理逻辑
*/
func login(rw http.ResponseWriter, request *http.Request) {
	// 获取请求的方法 request.Method 字符串类型的变量，返回GET, POST,PUT等method信息。
	fmt.Printf("method: %v\n", request.Method)
	//根据 request.Method 来判断是显示登录界面还是处理登录逻辑
	if request.Method == "GET" {
		crutime := time.Now().Unix() //获取经过的秒数
		h := md5.New()               //获取新的hash值
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprint("%x", h.Sum(nil)) //格式化
		//第一次登录网址时是显示登录界面，此时 request.Method 是 GET，后台机会使用template.ParseFiles 根据指定的文件创建模板示例（浏览器上的视图显示就是这么来的）。
		//使用template.ParseFiles 根据指定的文件创建模板示例
		t, _ := template.ParseFiles("web/login.html") //注意这个login.html的路径，否则会出现404错误
		//执行数据融合
		t.Execute(rw, token)
	}
	if request.Method == "POST" {
		//等我们输入用户名与密码再登录后 request.Method 就会成为POST。如果我们需要在服务器上输出传输过来的数据，就必须显式地调用 request.ParseForm()，默认情况下是不会显示调用的。
		//请求的是登录数据，那么执行登录的逻辑判断
		//解析传过来的参数，默认不会解析，必须显示调用后服务器才会输出参数信息
		request.ParseForm()
		//这里的request.Form["username"]可以用request.FormValue("username")代替，那么就不需要显示调用  request.ParseForm。但是如果同一个键中包含多个值，使用 request.FormValue 时只会返回第一个值
		//username := request.FormValue("username")
		//password := request.FormValue("password")
		//将字符串所有的空格消除后判断字符串username 、password是否为空字符串
		username := strings.Join(strings.Fields(strings.TrimSpace(request.FormValue("username"))), "")
		password := strings.Join(strings.Fields(strings.TrimSpace(request.FormValue("password"))), "")
		//template.HTMLEscapeString(username) 将username转义后 返回字符串
		fmt.Printf("username: %v\n", template.HTMLEscapeString(username)) //输出到服务器端
		fmt.Printf("password: %v\n", template.HTMLEscapeString(password))
		template.HTMLEscape(rw, []byte(username)) //输出到客户端
	}
}
