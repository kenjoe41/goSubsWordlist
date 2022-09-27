package cli

import (
	"os"
	"testing"
)

func TestCLI(t *testing.T) {
	if err := Cli(false, false); err == nil {
		t.Error("Expected error for no Stdin input")
	}

	userInput := "example.com\nsub.example.com\n\n"

	funcDefer, err := mockStdin(t, userInput)
	if err != nil {
		t.Errorf("%q", err)
	}

	defer funcDefer()
	if err := Cli(false, false); err != nil {
		t.Errorf("%q", err)
	}
}

// mockStdin is a helper function that lets the test pretend dummyInput as os.Stdin.
// It will return a function for `defer` to clean up after the test.
func mockStdin(t *testing.T, dummyInput string) (funcDefer func(), err error) {
	t.Helper()

	oldOsStdin := os.Stdin
	tmpfile, err := os.CreateTemp(t.TempDir(), t.Name())

	if err != nil {
		return nil, err
	}

	content := []byte(dummyInput)

	if _, err := tmpfile.Write(content); err != nil {
		return nil, err
	}

	if _, err := tmpfile.Seek(0, 0); err != nil {
		return nil, err
	}

	// Set stdin to the temp file
	os.Stdin = tmpfile

	return func() {
		// clean up
		os.Stdin = oldOsStdin
		os.Remove(tmpfile.Name())
	}, nil
}
