<a name="unreleased"></a>
## [Unreleased]


<a name="0.0.6"></a>
## [0.0.6] - 2019-10-18
### Added
- option `spinner.MessageEllipsis(string)`
- option `spinner.MaxMessageLength(int)`
- function `Truncate(string, int, interface{})`

### Feature
- `spinner.Message()` truncates message to predefined length

### Fixed
- unnecessary ellipsis


<a name="0.0.5"></a>
## [0.0.5] - 2019-10-18
### Changed
- `Interval(time.Duration)` option function


<a name="0.0.4"></a>
## 0.0.4 - 2019-10-16
### Added
- new char set `HalfClock2`
- option `spinner.CharSet([]string)`
- option `spinner.Variant(int)`
- option `spinner.ProgressIndicatorFormat(string)`
- order option
- `Reverse()` option
- `Prefix()` option
- new color set `C256YellowWhite`
- Changelog
- Example
- `Reversed` - spin in the opposite direction

### Fixed
- format artefacts
- internal package name `auxiliary`


[Unreleased]: https://github.com/alecrabbit/go-cli-spinner/compare/0.0.6...HEAD
[0.0.6]: https://github.com/alecrabbit/go-cli-spinner/compare/0.0.5...0.0.6
[0.0.5]: https://github.com/alecrabbit/go-cli-spinner/compare/0.0.4...0.0.5
