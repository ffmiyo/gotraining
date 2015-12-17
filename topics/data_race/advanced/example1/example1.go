// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// https://play.golang.org/p/MFiMZD_ODW

// Sample program to show a more complicated race condition using
// an interface value. This produces a read to an inteface value after
// a partial write.
package main

import "fmt"

// Speaker allows for speaking behavior.
type Speaker interface {
	Speak() bool
}

// Ben is a person who can speak.
type Ben struct {
	name string
}

// Speak allows Ben to say hello. It returns false if the method is
// called through the interface value after a partial write.
func (b *Ben) Speak() bool {
	if b.name != "Ben" {
		fmt.Printf("Ben says, \"Hello my name is %s\"\n", b.name)
		return false
	}

	return true
}

// Jerry is a person who can speak.
type Jerry struct {
	name string
}

// Speak allows Jerry to say hello. It returns false if the method is
// called through the interface value after a partial write.
func (j *Jerry) Speak() bool {
	if j.name != "Jerry" {
		fmt.Printf("Jerry says, \"Hello my name is %s\"\n", j.name)
		return false
	}

	return true
}

// main is the entry point for all Go programs.
func main() {
	// Create value of Ben and Jerry.
	ben := Ben{"Ben"}
	jerry := Jerry{"Jerry"}

	// Assign the value of type Ben to the interface value.
	person := Speaker(&ben)

	// Have a goroutine constantly assign the value of
	// type Ben to the interface.
	go func() {
		for {
			person = &ben
		}
	}()

	// Have a goroutine constantly assign the value of
	// type Jerry to the interface.
	go func() {
		for {
			person = &jerry
		}
	}()

	// Keep calling the Speak method against the interface
	// value until we have a race condition.
	for {
		if !person.Speak() {
			break
		}
	}
}
