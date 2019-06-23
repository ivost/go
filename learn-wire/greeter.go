package main

type Greeter struct {
	Message Message
}

func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}


func (g Greeter) Greet() Message {
	return g.Message
}

