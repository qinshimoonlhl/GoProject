package main

import (
	"fmt"
)

/*
定义接口
 */
type phone interface {
	call()
}

/*
定义结构体
*/
type NoKiaPhone struct {

}

/*
实现接口方法
*/
func (nokia NoKiaPhone) call(){
	fmt.Println("Nokia call U")
}

/*
定义结构体
 */
type iPhone struct {

}
/*
实现接口方法
*/
func (iphone iPhone) call(){
	fmt.Println("iPhone call U")
}

func main(){
	var pho phone
	pho = new (NoKiaPhone)
	pho.call()

	pho = new (iPhone)
	pho.call()


}