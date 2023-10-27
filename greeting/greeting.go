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

func generateRandomSalutation() string {
    salutations := []string{
        "Greetings, %v.",
        "Salutations, %v.",
        "Bienvenue, %v",
    }

    return salutations[rand.Intn(len(salutations))]
}
