package main

import (
	"log"
	"os"
	"flag"
//	"fmt"
)

func init() {
	log.SetOutput(os.Stdout)
}

func tryRunes() {
	// Go uses UTF-8 encoded characters in its strings.
	// Each character is of type rune.
	// This is a alias for int32 as a Unicode character
	// can be 1, 2 or 4 bytes in UTF-8 encoding.
	// In terms of characters, a string is an array of runes.
	str := "Здравейте 百度一下，你就知道"
	log.Println(str)
	log.Println("String length", len([]rune(str)))
	log.Println("Byte length", len(str))
}
	// cert, pem, tls
	// https://jan.newmarch.name/golang/security/chapter-security.html

func main() {
	flag.Parse()
	//tryNetutil()
	tryDbutil()
}

func tryDbutil() {
	TryGocql()
}

func tryNetutil() {

	host := "google.com"
	addr := Resolve(host)
	log.Printf("ip: %v", addr)

	hostPort := "google.com:80"
	resp := Head(hostPort)
	//log.Printf("head: %v", resp)

	url := "http://ip.jsontest.com/"
	resp = Get(url)
	log.Print(resp)

	url = "http://echo.jsontest.com/key1/value1/key2/value2"
	resp = Get(url)
	log.Print(resp)
	// parse JSON, key/value map
	fields := Parse(resp)
	for k, v := range fields {
		log.Printf("%s: %v", k, v)
	}

	// will exit
	//url = "http://foo.bar"
	//resp = Get(url)

	RenderTemplate()

	EchoServer()
}


