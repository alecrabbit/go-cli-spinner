# 🏵️  Go Console Spinner

[![Go Report Card](https://goreportcard.com/badge/github.com/alecrabbit/go-cli-spinner)](https://goreportcard.com/report/github.com/alecrabbit/go-cli-spinner)

```*** COMING SOON ***```
 
### Features todo list
- [x] progress indication during spin `spinner.Progress(0.5)` ➙ `50%`
- [x]  messages during spin `spinner.Message("message")`
- [x]  separated format settings for chars, messages and progress

    ```go
    spinner.FormatProgress = "[%4s]" // [  7%]
    ```
- [ ]  separated color settings for chars, messages and progress
- [ ]  has `Disable()` and `Enable()` methods (questionable)
- [x]  hides cursor on `spinner.Start()`, shows on `spinner.Stop()`
- [x]  cursor hide can be disabled `spinner.HideCursor = false` 
- [x]  has `Erase()` method
- [x]  final message `spinner.FinalMessage = "final message\n"`
- [x]  supports unix pipe `|` and redirect `>` output

It's a proof of concept and kinda port of [alecrabbit/php-console-spinner](https://github.com/alecrabbit/php-console-spinner)

> API is a subject to change

For now you can **try it as is** and shape it's development if you wish

> I'm developing it on Xterm terminal(package uses ANSI codes) so I hope it'll be fully functional in these environments. For other env's some help is required. Thank You.

> Works on Windows too! Thanks to [mattn/go-colorable](https://github.com/mattn/go-colorable)

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

    s := spinner.New(spinner.Snake2, 150*time.Millisecond)
    s.FinalMessage = "Done!\n"
    // s.HideCursor = false
    s.Reversed = true
    // s.Prefix = " " // spinner prefix
    s.FormatProgress = "[%4s]" // [  7%]

    rand.Seed(time.Now().UnixNano())
    // Start spinner
    s.Start()
    // for _, m := range messages {
    l := len(messages)
    for i, m := range messages {
        // Doing some work 1
        time.Sleep(500 * time.Millisecond)
        // Printing execution message
        {
            s.Erase() // optional if you're absolutely sure that your messages are longer
            fmt.Println(m)
            fmt.Print("..................................") // string to show that spinner can be used in inline mode
            s.Current() // Write current frame to output(optional - for smooth amination)
        }
        // Simulating spinner message
        if rand.Intn(16) > 7 {
            s.Message("") // Sometimes there are no messages
        } else {
            s.Message(fmt.Sprintf("Message at %s", time.Now().Format("15:04:05")))
        }
        // Doing some work 2
        time.Sleep(600 * time.Millisecond)
        // Simulating spinner progress
        s.Progress(float32(i) / float32(l)) // float32 0..1
    }
    time.Sleep(1 * time.Second)
    // Stop spinner
    s.Stop()
    time.Sleep(1 * time.Second)
}
``` 

> Also try to  redirect 
> ```
> go run main.go > out.txt
> ```
> and pipe 
> ```
> go run main.go | grep cess
> ```