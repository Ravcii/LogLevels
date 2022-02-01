package loglevels

import (
	"testing"
)

// defaultLevel for resetting it back after tests
const defaultLevel = DEBUG

func TestStringer(t *testing.T) {
	levels := []LoggingLevel{DEBUG, INFO, WARNING, ERROR}
	expected := []string{"Debug", "Info", "Warning", "Error"}

	for i, level := range levels {
		want := expected[i]
		got := level.String()

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	}

}

func TestSetDebugLevel(t *testing.T) {
	levels := []LoggingLevel{DEBUG, INFO, WARNING, ERROR}

	for _, want := range levels {
		SetLevel(want)
		got := currentLevel

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	}

	currentLevel = defaultLevel
}

func TestDefaultLevel(t *testing.T) {
	want := DEBUG
	got := currentLevel

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestGetLevel(t *testing.T) {
	levels := []LoggingLevel{DEBUG, INFO, WARNING, ERROR}

	for _, want := range levels {
		SetLevel(want)
		got := Level()

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	}

	currentLevel = defaultLevel
}

// func TestDebug(t *testing.T) {
// 	// Temporarily changing output so we can compare results
// 	log.Default().SetOutput()
// }
