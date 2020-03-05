package main

// type xxxx func 学习 and 接口
import "fmt"

type fun func(int)

type Call interface {
	call(int)
}

// callback 其实就是实现了 fun 这个type函数的具体实现
func callback(i int) {
	fmt.Println("i am callBack")
	fmt.Println(i)
}

//main中调用的函数  把 f 类型的函数作为参数给接受了进来
func one(i int, f func(int)) {
	two(i, fun(f))
}

//one()中调用的函数
func two(i int, c Call) {
	c.call(i)
}

//fun实现的Call接口的call()函数
func (f fun) call(i int) {
	f(i)
}

func main() {
	one(2, callback)
}
