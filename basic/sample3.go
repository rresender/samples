package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
)

func main() {

	// Print current directory
	currDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Current dir is %s\n", currDir)

	// Get current enviroment variable
	fmt.Printf("Current $PATH is %s\n", os.Getenv("PATH"))

	// Get current username
	currUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Current User name is %s", currUser.Username)

	// Get current username
	currUser, err = user.Current()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Current User name is %s\n", currUser.Username)

	//Note plataform dependent command
	cmd := exec.Command("ls", "-ltr")
	// Convenient wrapper
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", stdoutStderr)

}
