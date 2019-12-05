package main

import (
	"fmt"
	"math/rand"
	"time"
)

func apartnum(s []int, n int)  {
	//千位
	s[0] = n/1000
	//百位
	s[1] = n%1000/100
	//shiwei
	s[2] = n%1000%100/10
	//gewei
	s[3] = n%10
}

func gusstnum(s []int)  {
	var num int
	fmt.Println("输入一个四位数：")
	for {
		fmt.Scan(&num)
		if 999 < num && num < 10000 {
			break
		} else {
			fmt.Println("输入正确的四位数:")
			continue
		}
	}
	//调用apartnum分离
	inputslice := make([]int, 4)
	apartnum(inputslice,num)
	//判断
	n := 0
	for i := 0; i<4; i++ {
		if inputslice[i] > s[i] {
			fmt.Printf("第%d位数字猜大了！\n", i+1)
		} else if inputslice[i] < s[i] {
			fmt.Printf("第%d位数字猜小了\n", i+1)
		} else {
			fmt.Printf("第%d位数字猜对了！\n", i+1)
			n++
		}
	}
	if n == 4 {
		fmt.Println("恭喜你猜对了！")
	} else {
		fmt.Println("再来一次吧？")
		gusstnum(s)
		}

}



func main()  {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(9999)
	//保证为四位数
	for {
		if n > 999 {
			break
		}
	}
	randslice := make([]int, 4)
	fmt.Println(n)
//	将四位数n的位分离出来
	apartnum(randslice,n)
	//调用输入函数并判断
	gusstnum(randslice)


}