package main

import (
	"fmt"
	"time"

	"github.com/alecrabbit/go-cli-spinner"
)

func main() {
	s, _ := spinner.New()
	// Start spinner
	s.Start()

	// Doing some work
	fmt.Println("Message One: is written to StdOut.")
	time.Sleep(5 * time.Second)
	fmt.Println("Message Two: is written to StdOut.")

	// Stop spinner
	s.Stop()
}
