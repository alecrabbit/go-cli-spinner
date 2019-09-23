# 🏵️  Go Console Spinner

[![Go Report Card](https://goreportcard.com/badge/github.com/alecrabbit/go-cli-spinner)](https://goreportcard.com/report/github.com/alecrabbit/go-cli-spinner)

```*** COMING SOON ***```
 
It's a proof of concept, and kinda port of [alecrabbit/php-console-spinner](https://github.com/alecrabbit/php-console-spinner)

For now you can **try it as is** and shape it's development if you wish

> I'm developing it on Xterm terminal so I hope it'll be fully functional in these environments. For other env's some help is required. Thank You.


## Example

```go
package main

import (
    "fmt"
    "math/rand"
    "time"

    "github.com/alecrabbit/go-cli-spinner"
)

func main() {
    messages := []string{
        "Starting",
        "Initializing",
        "Gathering data",
        "Checking data",
        "Checking weather",
        "Processing",
        "Processing",
        "Processing",
        "Processing",
        "Processing",
        "Processing",
        "Processing",
        "Processing",
        "Almost there",
        "Be patient",
    }

    s := spinner.New(spinner.Snake3, 100*time.Millisecond)
    s.FinalMessage = "Done!\n"
    // s.HideCursor = false
    // s.Reversed = true

    rand.Seed(time.Now().UnixNano())
    s.Start()
    l := len(messages)
    for i, m := range messages {
        // Doing some work 1
        time.Sleep(500 * time.Millisecond)
        // Printing execution message
        {
            s.Erase()
            fmt.Println(m)
            s.Last()
        }
        // Simulating spinner message
        if rand.Intn(16) > 12 {
            s.Message("") // Sometimes there are no messages
        } else {
            s.Message(fmt.Sprintf("Message at %s", time.Now().Format("15:04:05")))
        }
        // Doing some work 2
        time.Sleep(500 * time.Millisecond)
        // Simulating spinner progress
        f := float32(i) / float32(l)
        s.Progress(f)
    }
    time.Sleep(3 * time.Second)
    s.Stop()
}
``` 
