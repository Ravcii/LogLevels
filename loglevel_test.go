package loglevels

import (
	"bytes"
	"io"
	"log"
	"testing"
)

// defaultLevel for resetting it back after tests
const defaultLevel = DEBUG

func TestStringer(t *testing.T) {
	levels := []LoggingLevel{DEBUG, INFO, WARNING, ERROR}
	expected := []string{"debug", "info", "warning", "error"}

	for i, level := range levels {
		want := expected[i]
		got := level.String()

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	}
}

func TestSetDebugLevel(t *testing.T) {
	defer SetLevel(defaultLevel)
	levels := []LoggingLevel{DEBUG, INFO, WARNING, ERROR}

	for _, want := range levels {
		SetLevel(want)
		got := currentLevel

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	}
}

func TestDefaultLevel(t *testing.T) {
	want := DEBUG
	got := currentLevel

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestGetLevel(t *testing.T) {
	defer SetLevel(defaultLevel)

	levels := []LoggingLevel{DEBUG, INFO, WARNING, ERROR}

	for _, want := range levels {
		SetLevel(want)
		got := Level()

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	}
}

// func TestDebug(t *testing.T) {
// 	// Temporarily changing output so we can compare results
// 	log.Default().SetOutput()
// }

func TestOutput(t *testing.T) {
	defer log.SetOutput(log.Default().Writer())
	defer log.SetFlags(log.Default().Flags())

	// TODO: clean this mess

	var buf = new(bytes.Buffer)
	log.SetOutput(buf)
	log.SetFlags(0)

	Output(ERROR, "test %s", "test2")
	want := "error:test test2\n"
	bytes, err := io.ReadAll(buf)
	if err != nil {
		t.Errorf("Can't read from buffer: %s", err.Error())
	}
	got := string(bytes)

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
