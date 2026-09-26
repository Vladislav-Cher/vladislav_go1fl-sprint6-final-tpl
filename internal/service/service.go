package service

import (
	"fmt"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// MorseTextConv определяет является ли переданная строка кодом Морзе или просто текстом,
// затем конвертирует текст в Морзе или наоборот
func MorseTextConv(data string) (string, error) {
	// символы которые могут встретиться в коде Морзе
	if data == "" {
		return "", fmt.Errorf("Передана пустая строка")
	}
	morseChar := ".- \n\r\t"

	for _, char := range data {
		if !strings.ContainsRune(morseChar, char) {
			return morse.ToMorse(data), nil
		}
	}

	return morse.ToText(data), nil
}
