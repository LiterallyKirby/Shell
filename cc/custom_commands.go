package cc

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"golang.org/x/crypto/ssh/terminal"
)

func Encrypt(input []string) {
	fmt.Println("Starting Encryption")
	if len(input) != 3 {
		fmt.Println("error: either too many or too few arguments")
		return
	}
	file_path := input[1]
	output_name := input[2]
	var password string

	for {
		fmt.Println("Enter Encryption Key: ")
		passwordBytes, err := terminal.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			fmt.Println("Key must be a string")
			return
		}

		fmt.Println("One more time for safety: ")
		password_checkBytes, err := terminal.ReadPassword(int(os.Stdin.Fd()))

		if err != nil {
			fmt.Println("Key must be a string")
			return
		}

		if string(passwordBytes) != string(password_checkBytes) {
			println("Passwords didn't match")
		} else {
			password = string(passwordBytes)
			break
		}
	}

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Failed to get current working directory: %v\n", err)
		return
	}
	fmt.Println("Current working directory:", cwd)

	// Reading file to encrypt
	fmt.Println("Reading File")
	to_encrypt, err := ioutil.ReadFile(file_path)
	if err != nil {
		fmt.Printf("failed to read file: %v\n", err)
		return
	}

	// Hashing Password
	hash := sha256.Sum256([]byte(password))
	hashed := hash[:] // Convert array to slice

	// Creating AES block
	fmt.Println("Creating AES Block")
	key := []byte(hashed) // Ensure the key is the correct length (16, 24, or 32 bytes)
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Printf("failed to create AES cipher: %v\n", err)
		return
	}

	// Creating GCM block
	fmt.Println("Creating GCM Block")
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Printf("failed to create GCM: %v\n", err)
		return
	}

	// Making a nonce
	fmt.Println("Making Nonce")
	nonce := make([]byte, gcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		fmt.Printf("failed to generate nonce: %v\n", err)
		return
	}

	// Encrypting the data
	fmt.Println("Encrypting Data")
	cipherText := gcm.Seal(nonce, nonce, to_encrypt, nil)

	// Writing the output file
	fmt.Println("Writing Output")
	if err := ioutil.WriteFile(output_name, cipherText, 0644); err != nil {
		fmt.Printf("failed to write encrypted file: %v\n", err)
		return
	}

	fmt.Println("Encryption successful. File saved:", output_name)

}

func Decrypt(input []string) {
	fmt.Println("Starting Decryption")
	if len(input) != 3 {
		fmt.Println("error: either too many or too few arguments")
		return
	}

	file_path := input[1]
	output_name := input[2]

	for {
		fmt.Print("Enter Decryption Key: ")
		passwordBytes, err := terminal.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			fmt.Printf("failed to read password: %v\n", err)
			return
		}
		fmt.Println() // Print a newline after the password input

		// Reading file to decrypt
		fmt.Println("Reading File")
		to_encrypt, err := ioutil.ReadFile(file_path)
		if err != nil {
			fmt.Printf("failed to read file: %v\n", err)
			return
		}

		// Hashing Password
		hash := sha256.Sum256([]byte(passwordBytes))
		hashed := hash[:] // Convert array to slice

		// Creating AES block
		fmt.Println("Creating AES Block")
		key := []byte(hashed) // Ensure the key is the correct length (16, 24, or 32 bytes)
		block, err := aes.NewCipher(key)
		if err != nil {
			fmt.Printf("failed to create AES cipher: %v\n", err)
			return
		}

		// Creating GCM block
		fmt.Println("Creating GCM Block")
		gcm, err := cipher.NewGCM(block)
		if err != nil {
			fmt.Printf("failed to create GCM: %v\n", err)
			return
		}

		// Extracting nonce
		nonceSize := gcm.NonceSize()
		if len(to_encrypt) < nonceSize {
			fmt.Println("Ciphertext is too short to contain nonce")
			return
		}
		nonce, cipherText := to_encrypt[:nonceSize], to_encrypt[nonceSize:]

		// Decrypting the data
		fmt.Println("Decrypting Data")
		plainText, err := gcm.Open(nil, nonce, cipherText, nil)
		if err != nil {
			fmt.Printf("failed to decrypt data: %v\n", err)
			fmt.Println("Invalid password, please try again.")
			continue // Invalid password; prompt again
		}

		// Writing the output file
		fmt.Println("Writing Output")
		if err := ioutil.WriteFile(output_name, plainText, 0644); err != nil {
			fmt.Printf("failed to write decrypted file: %v\n", err)
			return
		}

		fmt.Println("Decryption successful. File saved:", output_name)
		break // Exit the loop since decryption was successful
	}
}
func ChangeDir(command []string) {

	if err := os.Chdir(command[1]); err != nil {
		fmt.Printf("Error changing to directory %v\n", err)
	} else {
		wd, err := os.Getwd()
		if err != nil {
			fmt.Printf("failed to find wd, while the program will continue I would check that out %v\n", err)
		}
		fmt.Printf("Changed directory to %s\n", wd)
	}
}

func Find(command []string) {
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if strings.Contains(info.Name(), command[1]) {
			fmt.Println(path)
		}
		return nil
	})
}

func History(history []string) {
	for i, object := range history {
		i++
		istr := strconv.Itoa(i)
		fmt.Println(istr + "): " + object)
	}
}
