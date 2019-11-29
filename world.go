package main

import "fmt"

func Outer() func() int {	//返回值为一个函数
	var x int
	return func() int {
		x++
		return x * x		//变量x使用后直到下次使用都不会被重置
	}
}

func main()  {
	t := Outer()
	fmt.Println(t())
	fmt.Println(t())
	fmt.Println(t())
	fmt.Println(t())
	fmt.Println(t())
}
