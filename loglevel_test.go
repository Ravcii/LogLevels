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
	levels := []LogLevel{DEBUG, INFO, WARNING, ERROR}
	expected := []string{"debug", "info", "warning", "error"}

	for i, level := range levels {
		want := expected[i]
		got := level.String()

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	}
}

func logString(level LogLevel, format string, v ...interface{}) (string, error) {
	defer Logger().SetOutput(log.Default().Writer())
	var buf bytes.Buffer
	Logger().SetOutput(&buf)

	Output(level, format, v...)

	bytes, err := io.ReadAll(&buf)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func TestOutput(t *testing.T) {
	want := "debug:test1 test2\n"
	got, err := logString(DEBUG, "test1 %s", "test2")
	if err != nil {
		t.Errorf("can't read from test output:%s", err)
	}

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestDebugLevel(t *testing.T) {
	defer SetLevel(defaultLevel)
	SetLevel(DEBUG)

	levels := []LogLevel{DEBUG, INFO, WARNING, ERROR}
	shouldProduce := []string{
		"debug:debug output\n",
		"info:info output\n",
		"warning:warning output\n",
		"error:error output\n",
	}

	for i, level := range levels {
		want := shouldProduce[i]
		got, err := logString(level, "%s output", level.String())
		if err != nil {
			t.Error(err)
		}

		if got != want {
			t.Errorf("%d. got %q, wanted %q", i, got, want)
		}
	}
}

func TestInfoLevel(t *testing.T) {
	defer SetLevel(defaultLevel)
	SetLevel(INFO)

	levels := []LogLevel{DEBUG, INFO, WARNING, ERROR}
	shouldProduce := []string{
		"",
		"info:info output\n",
		"warning:warning output\n",
		"error:error output\n",
	}

	for i, level := range levels {
		want := shouldProduce[i]
		got, err := logString(level, "%s output", level.String())
		if err != nil {
			t.Error(err)
		}

		if got != want {
			t.Errorf("%d. got %q, wanted %q", i, got, want)
		}
	}
}

func TestWarningLevel(t *testing.T) {
	defer SetLevel(defaultLevel)
	SetLevel(WARNING)

	levels := []LogLevel{DEBUG, INFO, WARNING, ERROR}
	shouldProduce := []string{
		"",
		"",
		"warning:warning output\n",
		"error:error output\n",
	}

	for i, level := range levels {
		want := shouldProduce[i]
		got, err := logString(level, "%s output", level.String())
		if err != nil {
			t.Error(err)
		}

		if got != want {
			t.Errorf("%d. got %q, wanted %q", i, got, want)
		}
	}
}

func TestErrorLevel(t *testing.T) {
	defer SetLevel(defaultLevel)
	SetLevel(ERROR)

	levels := []LogLevel{DEBUG, INFO, WARNING, ERROR}
	shouldProduce := []string{
		"",
		"",
		"",
		"error:error output\n",
	}

	for i, level := range levels {
		want := shouldProduce[i]
		got, err := logString(level, "%s output", level.String())
		if err != nil {
			t.Error(err)
		}

		if got != want {
			t.Errorf("%d. got %q, wanted %q", i, got, want)
		}
	}
}
