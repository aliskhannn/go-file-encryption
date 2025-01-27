package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

func printMenu(actions map[int]string) {
	fmt.Print("\nChoose an action:\n\n")
	for k, v := range actions {
		fmt.Printf("%d. %s\n", k, v)
	}
}

func handleCreateFile() {
	fileName := getFileNameFromUser("\nEnter file name (example.txt):")
	if fileName == "" {
		return
	}

	fmt.Println("\nEnter text for the file:")
	text := getTextFromUser()

	err := createFile(fileName, text)

	if err != nil {
		fmt.Println("Error writing to file:", err)
	} else {
		fmt.Printf("File %s successfully created!\n", fileName)
	}
}

func handleEncryptFile() {
	fmt.Println("Encrypting a file...")
	// Реализация добавится позже
}

func handleDecryptFile() {
	fmt.Println("Decrypting a file...")
	// Реализация добавится позже
}

func getFileNameFromUser(prompt string) string {
	var fileName string

	fmt.Println(prompt)
	_, err := fmt.Scan(&fileName)

	if err != nil || !strings.HasSuffix(fileName, ".txt") {
		fmt.Println("Invalid file name. Please try again.")
		return ""
	}

	return fileName
}

func getTextFromUser() string {
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Invalid input. Please try again.")
		return ""
	}

	return text
}

func createFile(fileName, text string) error {
	err := os.WriteFile(fileName, []byte(text), 0644)
	if err != nil {
		return err
	}

	return nil
}
