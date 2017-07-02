package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Payload struct {
	Stuff Data
}

type Data struct {
	Fruit Fruits
	Veggies Vegetables
}

type Fruits map[string]int
type Vegetables map[string]int

func serve(w http.ResponseWriter, r *http.Request)   {
	resp, _ := makeResp()
	fmt.Fprint(w, string(resp))
}

func makeResp() ([]byte, error) {
	fruits := make(map[string]int)
	fruits["Apples"] = 25
	fruits["Oranges"] = 11

	veg := make(map[string]int)
	veg["Carrots"] = 10
	veg["Peppers"] = 0

	d := Data{fruits, veg}
	p := Payload{d}

	return json.MarshalIndent(p, "", "  ")

}

func main() {
	http.HandleFunc("/", serve)
	http.ListenAndServe("localhost:8282", nil)
}
