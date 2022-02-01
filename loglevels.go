package loglevels

import (
	"log"
	"strings"
)

// LoggingLevel is a syntax-sugar kind of type just for better readability
// for functions
type LoggingLevel int8

const (
	// DEBUG is the lowest logging level. Only devs should really care about it
	DEBUG LoggingLevel = iota
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

var currentLevel = DEBUG

func (l *LoggingLevel) String() string {
	return []string{"debug", "info", "warning", "error"}[*l]
}

// SetLevel changes
func SetLevel(level LoggingLevel) {
	currentLevel = level
}

// Level returns current logging level
func Level() LoggingLevel {
	return currentLevel
}

// Debug calls Output with a DEBUG logging level
func Debug(format string, v ...string) {
	if currentLevel < INFO {
		return
	}
	Output(DEBUG, format, v)
}

// Output sends formatted string to a default logger (without level checking!)
// In case you don't want to print any prefixes, set `level` to negative
func Output(level LoggingLevel, format string, v ...interface{}) {
	var sb strings.Builder
	if level >= DEBUG {
		sb.WriteString(level.String())
		sb.WriteByte(':')
	}
	sb.WriteString(format)
	log.Printf(sb.String(), v...)
}
