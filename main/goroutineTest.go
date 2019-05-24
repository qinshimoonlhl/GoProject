package main
import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func say(s string){
	for i := 0;i < 5;i++ {
		time.Sleep(100*time.Millisecond)
		fmt.Println(s)
	}
}

func main(){
	go say("hello")
	say("world")
	//缓冲区大小为2的通道
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2

	fmt.Println(<- ch)
	fmt.Println(<- ch)

	//var input int
	//fmt.Scanln(&input)//输入
	//fmt.Println(input)
	conn1,e := net.Dial("tcp","http://10.96.156.130:8899")
	fmt.Println(conn1)
	fmt.Println(e)

	r,err := http.Get("http://www.baidu.com")
	fmt.Println(ioutil.ReadAll(r.Body))
	fmt.Println(err)
}
