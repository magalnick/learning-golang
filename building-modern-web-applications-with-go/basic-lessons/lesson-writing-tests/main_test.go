package main

import "testing"

/**
 * The test files are located directly next to the files they're testing.
 * The name of the test file is the name of the module being tested, followed by "_test.go",
 * in this case, "main_test.go".
 * All test functions must start with "Test"
 * There must be a "go.mod" file. In this case I generated the file by running:
 * go mod init github.com/magalnick/learning-golang/building-modern-web-applications-with-go/lesson-writing-tests
 * To run the test, type in:
 * go test
 * add in the "-v" flag to make the output more verbose
 */

// simple tests to manually check the conditions one at a time
func TestDivide(t *testing.T) {
	_, err := divide(10, 1)
	if err != nil {
		t.Error("Got an error when we should not have")
	}
}

func TestBadDivide(t *testing.T) {
	_, err := divide(10, 0)
	if err == nil {
		t.Error("Did not get an error when one was expected")
	}
}

// more dynamic way of testing multiple scenarios at once
var tests = []struct {
	name     string
	dividend float64
	divisor  float64
	expected float64
	isErr    bool
}{
	{"valid-data", 100, 10, 10, false},
	{"invalid-data", 100, 0, 0, true},
}

func TestDivision(t *testing.T) {
	for _, tt := range tests {
		got, err := divide(tt.dividend, tt.divisor)
		if tt.isErr {
			if err == nil {
				t.Error("Expected an error but did not get one")
			}
		} else {
			if err != nil {
				t.Error("Did not expect an error but got one", err.Error())
			}
		}

		if got != tt.expected {
			t.Errorf("Expected %f but got %f", tt.expected, got)
		}
	}
}
