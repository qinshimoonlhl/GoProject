package main

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func uploadFile(c *gin.Context) {
	// FormFile方法会读取参数“upload”后面的文件名，返回值是一个File指针，和一个FileHeader指针，和一个err错误。
	file, header, err := c.Request.FormFile("upload")
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	// header调用Filename方法，就可以得到文件名
	filename := header.Filename
	fmt.Println(file, err, filename)

	// 创建一个文件，文件名为filename，这里的返回值out也是一个File指针
	out, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer out.Close()

	// 将file的内容拷贝到out
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}

	c.String(http.StatusCreated, "upload successful \n")
}

/*
func main() {
	router := gin.Default()

	// 调用POST方法，传入路由参数和路由函数
	//router.POST("/upload", uploadFile)
	//router.POST("/upload", uploadFile_single)
	router.POST("/upload", uploadMultiFile)
	// 监听端口8000，注意冒号。
	//router.Run(":8000")
}*/

func main() {
	//指定访问的路由
	http.HandleFunc("/upload", upload)
	//设定监听端口
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

/*
上传单个文件
*/
func uploadFile_single(c *gin.Context) {
	file, err := c.FormFile("upload")
	if err != nil {
		c.String(http.StatusBadRequest, "请求失败")
		return
	}
	//获取文件名
	fileName := file.Filename
	fmt.Println("文件名：", fileName)
	//保存文件到服务器本地
	if err := c.SaveUploadedFile(file, fileName); err != nil {
		c.String(http.StatusBadRequest, "保存失败 Error:%s", err.Error())
		return
	}
	c.String(http.StatusOK, "文件上传成功！")
}

/*
上传多个文件
*/
func uploadMultiFile(c *gin.Context) {
	//获取解析后表单
	form, _ := c.MultipartForm()
	files := form.File["upload[]"]
	//循环存文件到服务器本地
	for _, file := range files {
		c.SaveUploadedFile(file, file.Filename)
	}
	c.String(http.StatusOK, fmt.Sprintf("%d 个文件被上传成功"), len(files))
}

func upload(rw http.ResponseWriter, request *http.Request) {
	fmt.Println("method: ", request.Method)
	if request.Method == "GET" {
		//获取经过的秒数
		cruTime := time.Now().Unix()
		//获取新的hash值,防止页面重复提交
		h := md5.New()
		//将cruTime进行10进制转换，并写入io
		io.WriteString(h, strconv.FormatInt(cruTime, 10))
		//Sum将当前哈希附加到b并返回结果切片。
		token := fmt.Sprintf("%x", h.Sum(nil))
		//使用template.ParseFiles 根据指定的文件创建模板示例
		t, _ := template.ParseFiles("web/upload.html")
		//执行数据融合展现到客户端页面
		t.Execute(rw, token)
	} else {
		request.ParseMultipartForm(32 << 20)
		//通过 request.FormFile 获取客户端的文件句柄
		file, handler, err := request.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintln(rw, "%v", handler.Header)
		//根据指定的位置打开指定的文件
		fmt.Println("./test/" + handler.Filename)
		f, err := os.OpenFile("./main/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}
