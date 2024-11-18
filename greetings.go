package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// The `Hello` function in Go takes a name as input and returns a greeting message in Spanish welcoming
// the person by name.
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("nombre vacío")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, error := Hello(name)

		if error != nil {
			return nil, error
		}

		messages[name] = message
	}

	return messages, nil
}

func randomFormat() string {
	formats := []string{
		"¡Hola, %v!, ¡Bienvenido!",
		"¡Qué bueno verte, %v!",
		"¡Saludos, %v! ¡Encantado de conocerte!",
	}

	return formats[rand.Intn(len(formats))]
}
