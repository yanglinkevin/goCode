package main

import "fmt"

// 返回一个“返回int的函数”
func fibonacci() func() int {
	s1 := 0
	s2 := 1
	return func() int {
		//记录（back1）的值
		temp := s1

		// 重新赋值(这个就是核心代码)
		s1, s2 = s2, (s1 + s2)

		//返回temp
		return temp
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
