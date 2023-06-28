package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello 大写开头，代表public
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	// 异常场景LLT
	//message := fmt.Sprintf(randomFormat())
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

// randomFormat，小写开头，代表private 私有方法
func randomFormat() string {
	formats := []string{
		"Hi, %v, welcome!",
		"Great to see you, %v",
		"Hail, %v ! Well met!",
	}
	return formats[rand.Intn(len(formats))]
}
