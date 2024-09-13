package cli

import (
	"os"
	"testing"
)

func TestCLI_NoStdin(t *testing.T) {
	// Test when there is no input from Stdin
	if err := Cli(false, false); err == nil {
		t.Error("Expected error for no Stdin input")
	}
}

func TestCLI_WithStdin(t *testing.T) {
	// Test with simulated Stdin input
	userInput := "example.com\nsub.example.com\nanother.sub.example.com\n\n"

	funcDefer, err := mockStdin(t, userInput)
	if err != nil {
		t.Fatalf("Error mocking Stdin: %v", err)
	}
	defer funcDefer()

	if err := Cli(false, false); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestCLI_WithIncludeRoot(t *testing.T) {
	// Test including the root domain in the wordlist
	userInput := "example.com\nsub.example.com\n\n"

	funcDefer, err := mockStdin(t, userInput)
	if err != nil {
		t.Fatalf("Error mocking Stdin: %v", err)
	}
	defer funcDefer()

	if err := Cli(true, false); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestCLI_WithSilentFlag(t *testing.T) {
	// Test with the silent flag enabled (no banner output)
	userInput := "example.com\nsub.example.com\n\n"

	funcDefer, err := mockStdin(t, userInput)
	if err != nil {
		t.Fatalf("Error mocking Stdin: %v", err)
	}
	defer funcDefer()

	if err := Cli(false, true); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

// Test with invalid input
func TestCLI_InvalidInput(t *testing.T) {
	userInput := "invalid-url"

	funcDefer, err := mockStdin(t, userInput)
	if err != nil {
		t.Fatalf("Error mocking Stdin: %v", err)
	}
	defer funcDefer()

	if err := Cli(false, false); err != nil {
		t.Errorf("Unexpected error: %v", err)
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

	if _, err := tmpfile.Write([]byte(dummyInput)); err != nil {
		return nil, err
	}

	if _, err := tmpfile.Seek(0, 0); err != nil {
		return nil, err
	}

	// Set os.Stdin to the temp file
	os.Stdin = tmpfile

	return func() {
		// Clean up
		os.Stdin = oldOsStdin
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}, nil
}
