## Options

// TODO

> Note: some options shown below are not fully implemented yet

```go
    s, _ := spinner.New(
        // Set spinner variant
        spinner.Variant(spinner.Clock), // default spinner.Snake2
        // Override default refresh interval, each CharSet has it's own recommended refresh interval
        spinner.Interval(120),
        // Set your own character set
        spinner.CharSet([]string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}),
        // Override default color level support, default: TColor256
        spinner.ColorLevel(color.TColor256),
        // Override default elements order
        spinner.Order(spinner.Char, spinner.Progress, spinner.Message),
        // Override default progress element format
        spinner.ProgressFormat("%5s"),             // default: "%4s"
         // Override default progress indicator format
        spinner.ProgressIndicatorFormat("%.1f%%"), // 0.501 -> 50.1%, default: "%.0f%%" 0.501 -> 50%
        // Override default message format
        spinner.MessageFormat("(%s)"),   // (message)
        // Override default spinner element format
        spinner.Format("-%s -"),            // -⠏// -
        // Set prefix, default: ""
        spinner.Prefix("\x1b[38;5;161m>>\x1b[0m"),
        // Set final message, printed on s.Stop()
        spinner.FinalMessage("\x1b[38;5;34mDone!\x1b[0m\n"),
        // Spin in the opposite direction
        spinner.Reverse(),
        // Disable hide cursor 
        spinner.HideCursor(false),
    )
```

## Options order

- option `spinner.Interval(int)` should be after `spinner.Variant(int)`
- option `spinner.CharSet([]string)` should be after `spinner.Variant(int)`