package util

import (
	"strconv"
	"log"
)

func ToFloat64(s string) (res float64) {
	var err error
	res, err = strconv.ParseFloat(s, 64)
	if err != nil {
		log.Printf("Invalid number: %v", s)
		res = 0.0
	}
	return res
}
