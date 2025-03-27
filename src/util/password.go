package util

import (
	"math/rand"
)

const letters = "abcdefghijklmnopqrstuvwxyz"
const uppercaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numbers = "0123456789"
const symbols = "!@#$%^&*()_+-=[]{}|;:,.<>?~"

type GeneratePasswordOptions struct {
	IncludeUppercase bool
	IncludeNumbers   bool
	IncludeSymbols   bool
}

// generateRandomPassword generates a random password of the given length, excluding symbols
func GenerateRandomPassword(length int, opts GeneratePasswordOptions) string {
	bases := []string{letters}

	if opts.IncludeUppercase {
		bases = append(bases, uppercaseLetters)
	}
	if opts.IncludeNumbers {
		bases = append(bases, numbers)
	}
	if opts.IncludeSymbols {
		bases = append(bases, symbols)
	}

	return generatePassword(length, bases...)
}

func generatePassword(length int, bases ...string) (res string) {
	var acc = 0
	var sizes = make([]int, len(bases))
	for i := 0; i < len(bases)-1; i++ {
		l := length - acc - (len(bases) - i - 1)
		sizes[i] = rand.Intn(l)
		acc += sizes[i]
	}
	sizes[len(bases)-1] = length - acc

	for i := 0; i < len(bases); i++ {
		res += take(bases[i], sizes[i])
	}

	return shuffle(res)
}

func take(s string, size int) (res string) {
	for i := 0; i < size; i++ {
		index := rand.Intn(len(s))
		res += string(s[index])
	}
	return
}

func shuffle(s string) string {
	r := []rune(s)
	rand.Shuffle(len(r), func(i, j int) { r[i], r[j] = r[j], r[i] })
	return string(r)
}
