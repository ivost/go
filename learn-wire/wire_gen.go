// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

// Injectors from wire.go:

func InitializeEvent() (Event, error) {
	message := NewMessage()
	greeter := NewGreeter(message)
	event := NewEvent(greeter)
	return event, nil
}
