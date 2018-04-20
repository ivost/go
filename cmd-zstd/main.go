package main

import (
	"os/exec"
	"log"
	"fmt"
	"io"
	"strings"
	"bytes"
	"os"
)

func main() {
	test3()
}

func test3() {
	var out bytes.Buffer

	fp := os.Open("test.json")
}

func test1() {
	var out bytes.Buffer
	var err bytes.Buffer
	//cmd := exec.Command("tr", "a-z", "A-Z")
	//cmd := exec.Command("pwd")
	//cmd := exec.Command("zstd", "-f", "test.json")

	//cmd := exec.Command("zstd", "-f", "-o", "1.zst")
	cmd := exec.Command("zstd")

	cmd.Stdin = strings.NewReader("Hello World! 12345678901234567890")
	cmd.Stdout = &out
	cmd.Stderr = &err

	/////////////////
	check(cmd.Run())
	//\\\\\\\\\\\\\\\
	fmt.Printf("out: %v\n", out.String())
	fmt.Printf("err: %v\n", err.String())
}

func test2() {
	cmd := exec.Command("cat")
	stdin, err := cmd.StdinPipe()
	check(err)

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, "values written to stdin are passed to cmd's standard input")
	}()

	out, err := cmd.CombinedOutput()
	check(err)
	fmt.Printf("%s\n", out)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
