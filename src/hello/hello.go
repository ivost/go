package main

import (
		"fmt"
		"runtime"
		"time"
		)


func main() {
    fmt.Println("Hello")
    fmt.Println(time.Now().Format(time.RFC850))
    fmt.Println("GOOS", runtime.GOOS)
    fmt.Println("GOARCH", runtime.GOARCH)
}
