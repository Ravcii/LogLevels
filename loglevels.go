package loglevels

import (
	"log"
)

// LoggingLevel is a syntax-sugar kind of type just for better readability
// for functions
type LoggingLevel uint8

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
	return []string{"Debug", "Info", "Warning", "Error"}[*l]
}

// SetLevel changes
func SetLevel(level LoggingLevel) {
	currentLevel = level
}

// Level returns current logging level
func Level() LoggingLevel {
	return currentLevel
}

// Debug is a function for DEBUG logging level
func Debug(format string, v ...string) {
	if currentLevel < INFO {
		return
	}
	output(DEBUG, format, v)
}

// Info is a function for INFO logging level
func Info(format string, v ...string) {
	if currentLevel < INFO {
		return
	}
	log.Printf(format, v)
}

func output(level LoggingLevel, format string, v ...interface{}) {
	// var sb strings.Builder
	// TODO
}
