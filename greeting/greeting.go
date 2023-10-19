package greeting

import (
	"fmt"
)

func Greet(name string) string {
	message := fmt.Sprintf("Greetings, %v.", name)
	return message
}
