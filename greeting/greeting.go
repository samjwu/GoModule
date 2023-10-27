package greeting

import (
	"errors"
	"fmt"
)

func Greet(name string) (string, error) {
	if name == "" {
		return "", errors.New("name is empty")
    }

	message := fmt.Sprintf("Greetings, %v.", name)

	return message, nil
}
