package util

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strings"
)

const passwordLength = 64
const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// generateRandomPassword generates a random password of the given length, excluding symbols
func GenerateRandomPassword(length int) (string, error) {
	password := make([]byte, length)
	for i := range password {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		password[i] = letters[index.Int64()]
	}
	return string(password), nil
}

// receivePasswordFromStdin reads a password from stdin or generates a random password if none is provided
func ReceivePasswordFromStdin() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter password (or press Enter to generate a random one): ")
	input, _ := reader.ReadString('\n')
	password := strings.TrimSpace(input)

	// If no password is provided, generate one
	if password == "" {
		generatedPassword, err := GenerateRandomPassword(passwordLength)
		if err != nil {
			fmt.Println("Error generating password:", err)
			os.Exit(1)
		}
		fmt.Println("Generated password:", generatedPassword)
		password = generatedPassword
	}

	return password
}
