package main

import (
	"log"
)

//import "fmt"

func find(a []int, n int) int {

	lo := 0
	hi := len(a)

	for lo < hi {
		mid := lo + (hi - lo) / 2 	// avoid overflow
 		d := a[mid] - n
		//log.Printf("lo %d, hi %d, mid %d, a[mid] %d\n", lo, hi, mid, a[mid])
		if d == 0 {
			return mid
		}
		// lo----mid----hi
		//
		if d > 0 {
			hi = mid
		} else {
			lo = mid + 1
		}
	}

	return -1
}