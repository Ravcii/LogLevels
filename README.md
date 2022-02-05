# LogLevels

`LogLevels` is a simple package for implementing logging levels inspired by Python default logging package.

It's for those who think that using [Logrus]([https://github.com/sirupsen/logrus]) for their project just for splitting DEBUG/INFO logs is a bit of an overkill.
This package doesn't implement `Fatal` and `Panic` functions
## Installation

```
go get github.com/Ravcii/LogLevels
```

## Importing

```go
import log "github.com/Ravcii/LogLevels"
```

## Logging levels

There are 4 logging levels

* `log.DEBUG` is the lowest logging level. Only devs should really care about it
* `log.INFO` is a logging level for what's outputting what's happening right now in your code, but not as detailed as DEBUG
* `log.WARNING` is a logging level at which your program still runs but you want to at least wanna check what happened
* `log.ERROR` is a logging level at which something would likely crash your program

## Usage

### Levels
Use `SetLevel(level)` for settings a level.

Use `Level()` for getting a current level.

### Log Functions
There are 8 functions, 2 per each level for normal string and formatted string:
* `Debug()`, `Debugf()`
* `Info()`, `Infof()`
* `Warning()`, `Warningf()`
* `Error()`, `Errorf()`

 ### Changing default settings, changing writer, etc...

You can use `Logger()` to get default logger from the original go's `log` package.

Example:
```go
var buf bytes.Buffer

log.Logger().SetWriter(&buf)
```
