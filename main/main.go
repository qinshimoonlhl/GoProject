package main

import (
	"fmt"
)

func add(num1, num2 int) int {
	var sum int
	sum = num1 + num2
	return sum
}

func swap(x, y int) (int, int) {
	return y, x
}

//语言结构体
type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	const LENGH = 4
	var a string = "Golang"
	var b, c int = 1, 2
	var d = true
	e, f := 3, "hello"

	y, x := swap(b, c)
	if b < c {
		fmt.Println("b 小于 c")
	} else {
		fmt.Println("b 大于 c")
	}
	var (
		g int
		h string
	)
	var i, j, k int
	i, j, k = 7, 8, 9

	var grade string = "A"
	var score int = 100
	var n [5]int
	switch score {
	case 100:
		grade = "A++"
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 70:
		grade = "C"
	default:
		grade = "D"
	}
	for i = 0; i < 5; i++ {
		n[i] = i + 100
	}
	for j = 0; j < 5; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}
	fmt.Println(grade)
	fmt.Println(score)
	fmt.Println(b, c)
	fmt.Println(a)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Printf("变量的地址：%x\n", &k)
	fmt.Println(LENGH)
	fmt.Printf("加和是: %d\n", add(b, c))
	fmt.Println(y, x)

	var aa int = 20
	var pp *int

	pp = &aa
	fmt.Printf("aa 变量的地址是：%x\n", &aa)
	fmt.Printf("pp 变量的地址是：%x\n", pp)
	//结构体
	fmt.Println(Books{"Go语言", "lhl", "Go 语言教程", 7896578})

	//访问结构体
	var Books1 Books //声明Books类型的变量Books1
	Books1.author = "lhlhhj"
	Books1.book_id = 136554
	Books1.subject = "学习资料哦"
	Books1.title = "Python"

	fmt.Printf("Book1 title : %s\n", Books1.title)
	fmt.Printf("Book1 author : %s\n", Books1.author)
	fmt.Printf("Book1 subject : %s\n", Books1.subject)
	fmt.Printf("Book1 book_id : %d\n", Books1.book_id)

	//Go语言切片
	var number = make([]int, 2, 5)
	printSlice(number)
	if number == nil {
		fmt.Printf("空切片")
	}

	nums := []int{2, 3, 4}
	sums := 0
	for _, num := range nums {
		sums += num
	}
	fmt.Println("sunms:", sums)

	//GO语言集合
	var ProvenceMap map[string]string //创建map
	ProvenceMap = make(map[string]string)
	ProvenceMap["hubei"] = "wuhan"
	ProvenceMap["hunan"] = "changsha"
	ProvenceMap["jiangsu"] = "nanjing"

	for Provence := range ProvenceMap {
		fmt.Println(Provence, "省会是：", ProvenceMap[Provence])
	}

	prov, ok := ProvenceMap["henan"] /*判断是否存在*/
	if ok {
		fmt.Println("henan的省会是：", prov)
	} else {
		fmt.Println("不存在")
	}
	p, ok := ProvenceMap["jiangsu"]
	if ok {
		fmt.Println("henan的省会是：", p)
	}

	var m int = 13
	fmt.Printf("%d 的阶乘是 %d\n", m, Factorial(uint64(m)))

	testDefer()
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d sliec=%v\n", len(x), cap(x), x)
}

//递归函数
func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

func testDefer() {
	fmt.Println("----------------测试defer方法-----------------")
	i := 0
	defer fmt.Println(i) //输出0，因为i此时就是0
	i++
	defer fmt.Println(i) //输出1，因为i此时就是1
	return
}
