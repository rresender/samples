package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"reflect"
)

// SomeError sample
var SomeError = errors.New("error: description")

// EOFError sample
var EOFError = errors.New("EOF")

// CommandError struct
type CommandError struct {
	err string
}

func (c CommandError) Error() string {
	return c.err
}

// Avoid function
func Avoid() error {
	return &CommandError{"Avoid command"}
}

func main() {

	if EOFError == io.EOF {
		log.Fatal("Should no happen")
	}

	fmt.Println(reflect.TypeOf(SomeError))

	switch SomeError.(type) {
	case error:
		fmt.Println("Its an error")
	}

	switch SomeError {
	case SomeError:
		fmt.Println("Its the same error")
	}

	var invalidCommandError error = CommandError{"Invalid Command"}

	switch invalidCommandError.(type) {
	case CommandError:
		fmt.Println(invalidCommandError)
	}

	err := Avoid()

	if err, ok := err.(*CommandError); ok {
		fmt.Println(err)
	}

}
