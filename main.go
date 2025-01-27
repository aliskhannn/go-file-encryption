package main

import (
	"fmt"
)

func main() {
	var actions = map[int]string{
		1: "Create file",
		2: "Encrypt file",
		3: "Decrypt file",
		4: "Exit",
	}
	var cmd int

	fmt.Print("ROT13 File Encryption\n")

	for {
		printMenu(actions)

		_, err := fmt.Scan(&cmd)
		if err != nil {
			fmt.Println("Invalid input")
			return
		}

		switch cmd {
		case 1:
			handleCreateFile()
		case 2:
			handleEncryptFile()
		case 3:
			handleDecryptFile()
		case 4:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option, please try again.")
		}
	}
}

func handleCreateFile() {
	fmt.Println("Creating a file...")
	// Реализация добавится позже
}

func handleEncryptFile() {
	fmt.Println("Encrypting a file...")
	// Реализация добавится позже
}

func handleDecryptFile() {
	fmt.Println("Decrypting a file...")
	// Реализация добавится позже
}

func printMenu(actions map[int]string) {
	fmt.Print("\nChoose an action:\n\n")
	for k, v := range actions {
		fmt.Printf("%d. %s\n", k, v)
	}
}
