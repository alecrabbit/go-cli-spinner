# ```*** COMING SOON ***```
# 🏵️  Go Console Spinner

[![Go Report Card](https://goreportcard.com/badge/github.com/alecrabbit/go-cli-spinner)](https://goreportcard.com/report/github.com/alecrabbit/go-cli-spinner)

> API may be a subject to change

### Features todo list
- [x] highly configurable
- [x] progress indication during spin `spinner.Progress(0.5)` ➙ `50%`
- [x] messages during spin `spinner.Message("message")`
- [x] configurable elements order - chars, messages and progress
- [x] separated format settings for chars, messages and progress
- [ ] separated color settings for chars, messages and progress
- [ ] has `Disable()` and `Enable()` methods (questionable)
- [x] hides cursor on `spinner.Start()`, shows on `spinner.Stop()`
- [x] cursor hide can be disabled `spinner.HideCursor(false)` 
- [x] has `Erase()` method
- [x] has `Current()` method to write current frame again for smooth animation
- [x] final message
- [x] supports pipe `|` and redirect `>` output

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

### Usage

#### Method `spinner.Message`

```go
    // Set current message
    spinner.Message("Current message")
    // Hide message element
    spinner.Message("")
```

#### Method `spinner.Progress`

```go
    // Set current progress value 0..1
    spinner.Progress(0.511) // 51.1% 
    // Hide progress element
    spinner.Progress(0)
```
> Note: shown progress value depends on `ProgressIndicatorFormat`, default is "%0.f%%" and shows `70%`, for value `0.705`