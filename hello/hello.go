package main

import "fmt"
import "rsc.io/quote"
import "example.com/greetings"

func main() {
	fmt.Println("Hello world!")
	fmt.Println(quote.Go())
	messages := greetings.Hello("Gladys")
	fmt.Println(messages)
}
