package main

import (
	"fmt"
	"time"
)

func main1() {
	fmt.Println("Hello golang")

	//f := fib()
	//for i:=0; i<=64; i++ {
	//	fmt.Println(f())
	//}

	var a [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(a)

	s := a[0:3]

	fmt.Println(s)
	s = append(s, 9)
	fmt.Println(s)

	var x uint8 = 2
	y := 2
	fmt.Println(x)
	fmt.Println(y)

	go func(n int) { fmt.Printf("anonym gosub called with arg %v\n", n)} (42)

	// channel
	ch := make(chan bool)
	fmt.Println(Compare(2, 2));
	fmt.Println(Compare(40, 2));
	go gosub1(ch)
	ret := <- ch
	fmt.Printf("ret of gosub1 %t\n", ret)

	time.Sleep(1000 * time.Millisecond)

}

func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func Compare(a, b int) (bool, int) {
	if a > b {
		return false, a - b
	} else if a < b {
		return false, b - a
	}
	return true, 0
}

func gosub1(c chan bool) {
	fmt.Println("gosub1 called")
	c <- true
}