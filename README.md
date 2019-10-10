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

## [Examples](https://github.com/alecrabbit/go-cli-spinner/tree/master/examples/)
