package main

import (
	"bufio"

	"fmt"
	"os"
	"os/exec"
	"shell/cc"
	"strconv"
	"strings"

	"github.com/inancgumus/screen"
)

var cc_list = []string{"encrypt", "decrypt", "cd", "find", "history"}
var c_history = []string{}

func main() {
	screen.Clear()

	for {
		wd, err := os.Getwd()
		if err != nil {
			fmt.Printf("failed to find wd, while the program will continue I would check that out %v\n", err)
		}
		fmt.Print(wd, " $ ")
		reader := bufio.NewReader(os.Stdin)

		pre_command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Command must be a string")
			continue
		}

		// Trim newline characters
		pre_command = strings.TrimSpace(pre_command)

		// Detect special cases for history navigation
		if pre_command == "!!" && len(c_history) > 0 {
			// Repeat the last command
			pre_command = c_history[len(c_history)-1]
			fmt.Println(pre_command) // Show repeated command for user visibility
		} else if strings.HasPrefix(pre_command, "!") {
			// Run a specific history command by index
			numStr := pre_command[1:]
			num, err := strconv.Atoi(numStr)
			if err == nil && num > 0 && num <= len(c_history) {

				fmt.Printf("We are about to run the command: %s\n", c_history[num-1])
				fmt.Println("Are you SURE you want to do this? Y/N: ")
				reader := bufio.NewReader(os.Stdin)
				for {
					sure, err := reader.ReadString('\n')
					if err != nil {
						fmt.Println("Choice must be a string")
						continue
					}
					sure = strings.TrimSpace(sure)
					if sure == "y" || sure == "Y" {
						pre_command = c_history[num-1]
						break
					} else if sure == "n" || sure == "N" {
						fmt.Println("We won't run the command then")
						break
					} else {
						fmt.Println("That was not an option :/ try that again")
						continue
					}
				}

			} else {
				fmt.Println("No command found for that history index")
				continue
			}
		}

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
		if pre_command != "history" {
			c_history = append(c_history, pre_command)
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
	case "cd":
		cc.ChangeDir(command)
	case "find":
		cc.Find(command)
	case "history":
		cc.History(c_history)
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
