## Usage

### Method `spinner.Message`

Set current message
```go
spinner.Message("Current message")
```
Hide message element
```go
spinner.Message("")
```

#
### Method `spinner.Progress`

Set current progress value 0..1
```go
spinner.Progress(0.511) // 51.1% (depends on ProgressIndicatorFormat option)
```
Hide progress element
```go
spinner.Progress(0)
```
> Note: shown progress value depends on `ProgressIndicatorFormat` option, default is "%0.f%%" and shows `70%`, for value `0.705`