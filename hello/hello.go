package main

import (
	"fmt"
	"golang.org/x/example/stringutil"
	"log"
)
import "greetings"
import "upper"

func main() {
	log.SetPrefix("greeting: ")
	log.SetFlags(0)
	//messages, err := greetings.Hello("sky-we")
	//if err != nil {
	//	log.Fatal(err)
	//}
	names := []string{"ske-we1", "sky-we2", "sky-we3"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
	fmt.Println(stringutil.Reverse("Hello"))
	fmt.Println(upper.ToUpper("Hello"))
}
