package main

import (
	"fmt"
	"math/rand"
	"time"
)


func unirand(s []int)  {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i<len(s);i++ {
		s[i] = rand.Intn(100)
	}
}

func bullbu(s []int)  {
	n := len(s)
	for i := 0;i<n-1;i++ {
		for j :=0; j<n-1-i; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
}


func main() {
	n := 10
	s := make([]int, n)
	unirand(s)
	fmt.Println(s)
	//	排序
	bullbu(s)
	fmt.Println(s)

}