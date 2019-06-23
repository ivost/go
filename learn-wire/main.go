package main

import "log"

func main() {

	//message := NewMessage()
	//greeter := NewGreeter(message)
	//event := NewEvent(greeter)

	event, err := InitializeEvent()
	if err != nil {
		log.Fatal(err)
	}

	event.Start()
}


