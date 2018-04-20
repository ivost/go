package main

import (
	"os/exec"
	"log"
	"fmt"
	"io"
	"strings"
	"bytes"
	"os"
	"bufio"
)

func main() {
	deflate()
	inflate()
}

func deflate() {
	var stderr bytes.Buffer
	fname := "test.json"
	fp, err := os.Open(fname)
	check(err)
	defer fp.Close()
	fq, err := os.Create(fname+".zst")
	check(err)
	defer fq.Close()

	cmd := exec.Command("zstd", "-f")
	r := bufio.NewReader(fp)
	cmd.Stdin = r
	w := bufio.NewWriter(fq)
	cmd.Stdout = w
	check(cmd.Run())
	w.Flush()
	fmt.Printf("stderr: %v\n", stderr.String())
}

func inflate() {
	var stderr bytes.Buffer
	fname := "test.json"
	fp, err := os.Open(fname+".zst")
	defer fp.Close()
	check(err)
	fq, err := os.Create("1"+fname)
	check(err)
	defer fp.Close()
	cmd := exec.Command("zstd", "-f", "-d")
	r := bufio.NewReader(fp)
	cmd.Stdin = r
	w := bufio.NewWriter(fq)
	cmd.Stdout = w
	cmd.Stderr = &stderr
	check(cmd.Run())
	w.Flush()
	fmt.Printf("stderr: %v\n", stderr.String())
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
