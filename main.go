package main

import (
	"bufio"

	"fmt"
	"os"
	"os/exec"
	"shell/cc"
	"strings"

	"github.com/inancgumus/screen"
)

var cc_list = []string{"encrypt", "decrypt"}

func main() {
	screen.Clear()
	for {

		fmt.Print("$ ")
		reader := bufio.NewReader(os.Stdin)

		pre_command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Command must be a string")
			continue
		}

		// Trim newline characters
		pre_command = strings.TrimSpace(pre_command)

		if pre_command == "" {
			fmt.Println("Command can't be empty")
			continue
		}

		// Split the command into parts
		command := strings.Split(pre_command, " ")

		// Check if the command exists in the list
		if contains(cc_list, command[0]) {
			run_cc(command)
		} else {
			run_command(command)
		}
	}
}

// Helper function to check if an item is in a slice
func contains(slice []string, item string) bool {
	for _, element := range slice {
		if element == item {
			return true
		}
	}
	return false
}

// figure out which command to run
func run_cc(command []string) {
	switch command[0] {
	case "encrypt":
		cc.Encrypt(command)
	case "decrypt":
		cc.Decrypt(command)
	}
}

func run_command(command []string) {
	// command[0] is the program to run, and command[1:] are the arguments
	cmd := exec.Command(command[0], command[1:]...)

	// Redirect output to the terminal
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	// Run the command
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error executing command: %v\n", err)
	}
}
