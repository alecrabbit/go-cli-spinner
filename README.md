# ```*** COMING SOON ***```
# 🏵️  Go Console Spinner

[![Go Report Card](https://goreportcard.com/badge/github.com/alecrabbit/go-cli-spinner)](https://goreportcard.com/report/github.com/alecrabbit/go-cli-spinner)

> API may be a subject to change

### Features
- highly configurable ([options](docs/options.md))
- progress indication during spin `spinner.Progress(0.5)` ➙ `50%`
- messages during spin `spinner.Message("message")`
- configurable elements order - chars, messages and progress
- separated format settings for chars, messages and progress
- hides cursor on `spinner.Start()`, shows on `spinner.Stop()`
- cursor hide can be disabled `spinner.HideCursor(false)` 
- has `Erase()` method
- has `Current()` method to write current frame again for smooth animation
- final message
- supports pipe `|` and redirect `>` output

- [ ] separated color settings for chars, messages and progress
- [ ] has `Disable()` and `Enable()` methods (questionable)

It's a proof of concept and kinda port of [alecrabbit/php-console-spinner](https://github.com/alecrabbit/php-console-spinner)

For now you can try it **as is** and shape it's development if you wish

> Xterm terminal(package uses ANSI codes) 

> Works on Windows too! Thanks to [mattn/go-colorable](https://github.com/mattn/go-colorable)

### [Examples](https://github.com/alecrabbit/go-cli-spinner/tree/master/examples/)

### Quickstart

```go
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
    time.Sleep(3 * time.Second)
    // Set current message
    s.Message("Current message")
    // Set current progress
    s.Progress(0.511)
    // Doing some work
    time.Sleep(3 * time.Second)
    // Stop spinner
    s.Stop()
}
```

### [Usage](docs/usage.md)