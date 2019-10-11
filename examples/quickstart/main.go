package main

import (
	"time"

	"github.com/alecrabbit/go-cli-spinner"
)

func main() {
	s, _ := spinner.New()
	// Start spinner
	s.Start()
	// Doing some work
	time.Sleep(10 * time.Second)
	// Stop spinner
	s.Stop()
}
