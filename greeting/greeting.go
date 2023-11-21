package greeting

import (
	"errors"
	"fmt"
	"math/rand"
)

func Greet(name string) (string, error) {
	if name == "" {
		return "", errors.New("Error: name is empty")
    }

	message := fmt.Sprintf(generateRandomSalutation(), name)

	return message, nil
}

func GreetAll(names[] string) (map[string]string, error) {
	messages := make(map[string]string)
    
	for _, name := range names {
        message, err := Greet(name)

        if err != nil {
            return nil, err
        }
		
        messages[name] = message
    }

    return messages, nil
}

func generateRandomSalutation() string {
    salutations := []string{
        "Greetings, %v.",
        "Salutations, %v.",
        "Bienvenue, %v",
    }

    return salutations[rand.Intn(len(salutations))]
}
