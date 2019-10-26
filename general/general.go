package main

import (
	"fmt"
	"math"
	"sync"
)

// variable declarations
var atom = 1

const monad = "monad"

// know when init is called
func init() {
	atom = 1
}

// functions declarations
func calculateCeil(input float64) (result float64, err error) {
	return math.Ceil(input), nil
}

// type definitions and aliases

// ByteSlice Type
type ByteSlice []byte

// Bytes Type
type Bytes ByteSlice

// main
func main() {
	//some basic printf flags
	a := fmt.Sprintf("%d\n", atom)
	b := fmt.Sprintf("%s\n", monad)
	fmt.Printf("%s%s", a, b)

	// short declaration assignment
	stringSlice := [3]string{"1", "2", "3"}

	// basic slice operations
	fmt.Println(len(stringSlice))

	// underscore for ignoring results
	_, err := calculateCeil(1.3)

	// Checking for errors
	if err != nil {
		panic("Oops")
	}

	// infinite loop
	counter := 0
	for {
		if counter > 1000 {
			break
		} else {
			counter++
		}
	}

	// type switches
	var boolRef interface{}
	boolean := true
	boolRef = &boolean

	switch t := boolRef.(type) {
	default:
		fmt.Printf("Unexpected TYpe %T\n", t)
	case bool:
		fmt.Printf("Boolean %t\n", t)
	case int:
		fmt.Printf("Interger %d\n", t)
	case *bool:
		fmt.Printf("Pointer to Boolean %t\n", *t)
	case *int:
		fmt.Printf("Pointer to Interger %d\n", *t)
	}

	// defer calls as lifo order
	defer fmt.Printf("%d", 1)
	defer fmt.Printf("%d", 2)

	l := new(sync.Mutex)
	l.Lock()
	defer l.Unlock()

	// Interface and export rules
	type Iterator interface {
		Next() interface{}
		hasNext() bool
	}

	type Guard struct{}

	// Channels
	ch := make(chan Guard, 1)

	// Goroutines
	go func(chan Guard) {
		fmt.Printf("%T", ch)
		ch <- Guard{}
		close(ch)
	}(ch)

	<-ch
}
