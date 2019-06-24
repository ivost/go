package main

import "log"

/*
go run main.go greeter.go event.go message.go wire_gen.go
 */

func main() {
	event, err := InitializeEvent()
	if err != nil {
		log.Fatal(err)
	}

	event.Start()
}


