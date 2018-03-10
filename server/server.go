package main

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"runtime/pprof"
	"strings"

	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

var cpuprof = "cpu.prof"
var memprof = "mem.prof"

func main() {
	router := routing.New()

	router.Get("/foo", func(c *routing.Context) error {
		write1(c)
		return nil
	})

	panic(fasthttp.ListenAndServe(":8080", router.HandleRequest))
}

var line1 string

func init() {
	line1 = strings.Repeat("# xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx\n", 13)
	line1 = line1 + "\n# ~1KB of junk to force browsers to start rendering immediately: \n"
}

func write1(c *routing.Context) {
	//c.Request.Header().Set("Content-Type", "text/plain")
	io.WriteString(c, line1)
	//w.(http.Flusher).Flush()
}

func profileCpu() bool {
	if cpuprof == "" {
		return false
	}

	f, err := os.Create(cpuprof)
	if err != nil {
		fmt.Printf("Error creating cpu.prof: %v\n", err)
		return false
	}
	fmt.Printf("StartCPUProfile %s\n", cpuprof)
	if err := pprof.StartCPUProfile(f); err != nil {
		fmt.Printf("could not start CPU profile: %v\n", err)
		return false
	}
	return true
}

func profileMem() bool {
	if memprof == "" {
		return false
	}
	f, err := os.Create(memprof)
	if err != nil {
		fmt.Printf("could not create memory profile: %v\n", err)
		return false
	}
	fmt.Printf("StartHeapProfile %s\n", memprof)
	runtime.GC() // get up-to-date statistics
	if err := pprof.WriteHeapProfile(f); err != nil {
		fmt.Printf("could not write memory profile: %v\n", err)
		return false
	}
	f.Close()
	return true
}
