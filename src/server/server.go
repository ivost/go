package main

import (
	"fmt"
	"log"
	"net/http"
	//"sync"
	//"strconv"
	"html/template"
	"os"
	"io"
	//"time"
	//"crypto/md5"
	//"strconv"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Printf("%s\n", r.RequestURI)
		file := r.URL.Path[1:]
		if file == "/" {
			file = "index.html"
		}
		http.ServeFile(w, r, file)
	})

	http.HandleFunc("/login", login)
	http.HandleFunc("/upload", upload)

	log.Fatal(http.ListenAndServe(":8081", nil))

}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login - method:", r.Method)
	if r.Method == "GET" {
		//t, _ := template.ParseFiles("index.html")
		//t.Execute(w, nil)
	} else {
		r.ParseForm()
		user := template.HTMLEscapeString(r.Form.Get("user"))
		pass := template.HTMLEscapeString(r.Form.Get("pass"))
		fmt.Fprintf(w, "User: %s", user)
		fmt.Println("username:", user)
		fmt.Println("password:", pass)
	}
}

func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("upload - method:", r.Method)
	if r.Method != "POST" {
		return
	}

	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	//fmt.Fprintf(w, "%v", handler.Header)
	fmt.Fprintf(w, "File uploaded")

	f, err := os.OpenFile("./data/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
}

//var counter int
//var mutex = &sync.Mutex{}
//func echoString(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "hello")
//}
//
//func incrementCounter(w http.ResponseWriter, r *http.Request) {
//	mutex.Lock()
//	counter++
//	fmt.Fprintf(w, strconv.Itoa(counter))
//	mutex.Unlock()
//}
