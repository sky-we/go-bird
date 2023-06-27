package greetings

import "fmt"

// Hello 大写开头，其他可调用
func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
