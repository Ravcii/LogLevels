# LogLevels

`LogLevels` is a simple package for implementing logging levels inspired by Python default logging package.

It's for those who think that importing [Logrus]([https://github.com/sirupsen/logrus]) to their project just for splitting DEBUG/INFO logs is a bit of an overkill.

## Installation

TODO

## Logging levels

There are 6 logging levels (4 from Python, 2 Go-specific)

* `DEBUG` - Debug leve
* `INFO` - Info level
* `WARNING` - Warning level
* `ERROR` - Error level
* `FATAL` - Go-specific FATAL level, calls `os.exit(1)` after printing
* `PANIC` - Go-specific FATAL level, calls `panic()` after printing