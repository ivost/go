package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"io/ioutil"
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


func main() {
	url := "http://localhost:8282"
	res, _ := http.Get(url)
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	var p Payload

	json.Unmarshal(body, &p)

	fmt.Println(p.Stuff.Fruit, "\n", p.Stuff.Veggies)

}
