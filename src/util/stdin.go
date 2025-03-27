package util

import (
	"bufio"
	"fmt"
	"os"

	"syscall"

	"golang.org/x/term"
)

func ReadStdin() []byte {
	// Read from stdin
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Bytes()

	// Handle stdin error
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading stdin:", err)
		os.Exit(1)
	}

	return input
}

func PromptPassword() string {
	fmt.Print("Enter password: ")
	password, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		panic(err)
	}
	return string(password)
}

func PromptPasswordConfirm() string {
	fmt.Print("Enter new password: ")
	password1, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		panic(err)
	}
	fmt.Println()

	fmt.Print("Confirm password: ")
	password2, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		panic(err)
	}
	fmt.Println()

	if string(password1) != string(password2) {
		panic("passwords do not match")
	}

	return string(password1)
}
