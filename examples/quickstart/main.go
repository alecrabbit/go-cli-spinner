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
    {
        time.Sleep(1 * time.Second)
        s.Progress(0.1) // -> 10%
        fmt.Println("Message One: is written to StdOut.")
        time.Sleep(1 * time.Second)
        s.Progress(0)        // Hide progress indicator
        s.Message("Message") // Set message
        time.Sleep(1 * time.Second)
        fmt.Println("Message Two: is written to StdOut.")
        s.Message("") // Hide message
        time.Sleep(1 * time.Second)
        s.Progress(1.0) // -> 100%
        time.Sleep(500 * time.Millisecond)
    }

    // Stop spinner
    s.Stop()
}
