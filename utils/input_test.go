package utils

import (
	"testing"
)

func TestGetSessionToken(t *testing.T) {
	_, err := getSessionToken()
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
}

func TestGetInput(t *testing.T) {
    day := 1

    input1, err := GetInput(day)
    if err != nil {
        t.Fatalf("First call failed: %v", err)
    }

    if len(input1) == 0 {
        t.Errorf("Expected non-empty input")
    }

    path := getInputPath(day)
    if !fileExists(path) {
        t.Errorf("Cache file was not created")
    }

    input2, err := GetInput(day)
    if err != nil {
        t.Fatalf("Second call failed: %v", err)
    }  

    if input1 != input2 {
        t.Errorf("Cached content differs from downloaded content")
    }

    t.Logf("Successfully fetched and cached input for day %d", day)
}