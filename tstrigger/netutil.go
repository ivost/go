package main

import (
	"gopkg.in/resty.v0"
	"net"
	"fmt"
	"os"
	"io/ioutil"
	"log"
//	"io"
	"bytes"
	"io"
	"encoding/json"
	"strings"
	"html/template"
)


func Get(url string) string {
	//https://github.com/go-resty/resty

	// GET request
	resp, err := resty.R().Get(url)
	checkError(err)

	// explore response object

	//if err != nil {
	//	log.Printf("Error: %v", err)
	//	return ""
	//}
	//log.Printf("Response Status Code: %v", resp.StatusCode())
	//log.Printf("Response Status: %v", resp.Status())
	//log.Printf("Response Time: %v", resp.Time())
	//log.Printf("Response Recevied At: %v", resp.ReceivedAt())
	//log.Printf("Response Body: %v", resp.String())
	return resp.String()
}

// parse json string, return map key/value
// https://talks.golang.org/2015/json.slide#5
func Parse(s string) map[string]string {
	fields := map[string]string{}

	dec := json.NewDecoder(strings.NewReader(s))
	if err := dec.Decode(&fields); err != nil {
		log.Printf("decode error: %v in %s", err, s)
		return nil
	}
	return fields
}

func Resolve(host string) string {
	addr, err := net.ResolveIPAddr("ip", host)
	checkError(err)
	return addr.String()
}

func Head(hostPort string) string {
	addr, err := net.ResolveTCPAddr("tcp4", hostPort)
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, addr)
	checkError(err)

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	//result, err := readFully(conn)
	result, err := ioutil.ReadAll(conn)
	checkError(err)

	return string(result)
}

func EchoServer() {
	port := 1201
	log.Printf("EchoServer listening on port %v" ,port)
	service := fmt.Sprintf(":%d", port)
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	//listener, err := net.Listen("tcp", service)
	//checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		// run as a goroutine
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	// close connection on exit
	defer conn.Close()
	log.Printf("Handle connection %v", conn.RemoteAddr().String())
	var result [512]byte
	for {
		//_, err := io.Copy(conn, conn)
		//result, err := readFully(conn)
		n, err := conn.Read(result[0:])
		if err != nil {
			log.Printf("read error %v", err)
			return
		}
		//log.Printf("result %s", string(result[0:n]));
		_, err2 := conn.Write(result[0:n])
		if err2 != nil {
			log.Printf("write error %v", err2)
			return
		}
	}
}

func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close()

	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return result.Bytes(), nil
}


type Person struct {
	Name   string
	Age    int
	Emails []string
	Jobs   []*Job
}

type Job struct {
	Employer string
	Role     string
}

const templ = `The name is {{.Name}}.
The age is {{.Age}}.
{{range .Emails}}
        An email is {{.}}
{{end}}

{{with .Jobs}}
    {{range .}}
        An employer is {{.Employer}}
        and the role is {{.Role}}
    {{end}}
{{end}}
`

func RenderTemplate() {
	job1 := Job{Employer: "Monash", Role: "Honorary"}
	job2 := Job{Employer: "Box Hill", Role: "Head of HE"}

	person := Person{
		Name:   "jan",
		Age:    50,
		Emails: []string{"jan@newmarch.name", "jan.newmarch@gmail.com"},
		Jobs:   []*Job{&job1, &job2},
	}

	t := template.New("Person template")
	t, err := t.Parse(templ)
	checkError(err)

	err = t.Execute(os.Stdout, person)
	checkError(err)
}


func checkError(err error) {
	if err != nil {
		log.Fatalf("Fatal error: %s", err.Error())
		os.Exit(1)
	}
}