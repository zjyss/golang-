package main

import "fmt"

type Student struct {
	name string
	id int
	age int
}

func (p Student) Setinfovalue()  {
	fmt.Printf("Setincfovalue : %p , %v\n", &p, p)
}

func (p *Student) Setinfopointer()  {
	fmt.Printf("Setinfopointer : %p , %v\n", p, p)
}

func main()  {
	p := Student{"zhang", 1, 25}
	fmt.Printf("main : %p , %v\n", &p, p)
	pFunc1 := p.Setinfovalue
	pFunc1()
	pFunc2 := p.Setinfopointer
	pFunc2()

	//方法表达式
	f := (*Student).Setinfopointer
	f(&p)
	f2 := (Student).Setinfovalue
	f2(p)
}