package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"server/store"
	"os"
	"flag"
)

func init() {
	log.SetOutput(os.Stdout)
	flag.Parse()
}

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/login", login)
	http.HandleFunc("/upload", store.Upload)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func root(w http.ResponseWriter, r *http.Request) {
	//fmt.Printf("%s\n", r.RequestURI)
	file := r.URL.Path[1:]
	if file == "/" {
		file = "index.html"
	}
	http.ServeFile(w, r, file)
}

func login(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("login - method:", r.Method)
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

