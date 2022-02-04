package loglevels

import (
	"log"
	"strings"
)

// LogLevel is a syntax-sugar kind of type just for better readability
// for functions
type LogLevel int8

const (
	// DEBUG is the lowest logging level. Only devs should really care about it
	DEBUG LogLevel = iota
	// INFO is a logging level for what's outputting what's happening right now
	// in your code, but not as detailed as DEBUG
	INFO
	// WARNING is a logging level at which your program still runs but you want
	// to at least have a look what happened
	WARNING
	// ERROR is a logging level at which something would likely crash your program
	// Default Go `Fatal` and `Panic` log calls also falls in this category
	ERROR
)

var logger = log.New(log.Default().Writer(), "", 0)
var currentLevel = DEBUG

func (l *LogLevel) String() string {
	return []string{"debug", "info", "warning", "error"}[*l]
}

// Logger returns default logger for changing settings
func Logger() *log.Logger {
	return logger
}

// SetLevel changes the minimal level required for output
func SetLevel(level LogLevel) {
	currentLevel = level
}

// Level returns current logging level
func Level() LogLevel {
	return currentLevel
}

// Debug outputs a debug-level log
func Debug(str string) {
	Debugf(str, nil)
}

// Debugf outputs a formatted debug-level log
func Debugf(format string, v ...interface{}) {
	Output(DEBUG, format, v)
}

// Info outputs an info-level log
func Info(str string) {
	Infof(str, nil)
}

// Infof outputs a formatted info-level log
func Infof(format string, v ...interface{}) {
	Output(INFO, format, v)
}

// Warning outputs a warning-level log
func Warning(str string) {
	Warningf(str, nil)
}

// Warningf outputs a formatted warning-level log
func Warningf(format string, v ...interface{}) {
	Output(WARNING, format, v)
}

// Error outputs an error-level log
func Error(str string) {
	Errorf(str, nil)
}

// Errorf outputs a formatted error-level log
func Errorf(format string, v ...interface{}) {
	Output(ERROR, format, v)
}

// Output sends formatted string to a default logger (without level checking!)
// In case you don't want to print any prefixes, set `level` to negative
func Output(level LogLevel, format string, v ...interface{}) {
	var sb strings.Builder
	if level >= currentLevel {
		sb.WriteString(level.String())
		sb.WriteString(":")
	}
	sb.WriteString(format)

	logger.Printf(sb.String(), v...)
}
