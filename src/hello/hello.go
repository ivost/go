package main

import (
		"fmt"
		"time"
		)

func main() {
    fmt.Printf("hello, world\n")
    fmt.Println(time.Now().Format(time.RFC850))
}
