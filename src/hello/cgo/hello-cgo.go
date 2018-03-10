package main

/*
  #include <stdlib.h>
  #include <event2/event.h>
  #cgo LDFLAGS: -levent
 */
import "C"
import (
		"fmt"
		"os"
	  	"unsafe"
		"time"
		)

// testing cgo
// https://blog.gopheracademy.com/advent-2015/libc-hooking-go-shared-libraries/
// https://github.com/golang/go/blob/master/misc/cgo/testshared/shared_test.go#L416

func main() {
	const key = "FOO"
	const val = "BAR"

    fmt.Println("Testing go - C calls")
    fmt.Println(time.Now().Format(time.RFC850))
    // set env.var (go)
  	fmt.Printf("os.Setenv(%q) = %q\n", key, val)
  	os.Setenv(key, val)

  	// read back 
  	// key is go string - must convert to C string
  	// and free memory allocated by C
  	keyc := C.CString(key)
  	defer C.free(unsafe.Pointer(keyc))
  	v := C.getenv(keyc)
  	if uintptr(unsafe.Pointer(v)) == 0 {
  		fmt.Println("getenv returned NULL")
  		return
  	}
  	// result is C string - musyt convert to go string
  	vs := C.GoString(v)
  	fmt.Printf("C.getenv(%q) = %q\n", key, vs)
  	if vs != val {
  		fmt.Printf("getenv() = %q; want %q\n", vs, val)
  	}

  	v = C.event_get_version()
  	fmt.Printf("libevent version: %q\n", C.GoString(v))

}
