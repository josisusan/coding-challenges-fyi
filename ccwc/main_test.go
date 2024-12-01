package main

import (
	"strings"
	"testing"
)

func TestCountTotalBytes(t *testing.T) {
	t.Run("returns the total bytes in a reader", func(t *testing.T) {
		expectedTotal := 67
		total, _ := countTotalBytes(strings.NewReader(`Hello, this is my test file.
To Test the line count.
It should be 3`))

		if total != expectedTotal {
			t.Errorf("Expected: %d, Got: %d", expectedTotal, total)
		}
	})
}

func TestCountLines(t *testing.T) {
	t.Run("returns the total lines in a reader", func(t *testing.T) {
		expectedTotal := 3
		total, _ := countLines(strings.NewReader(`Hello, this is my test file.
To Test the line count.
It should be 3`))

		if total != expectedTotal {
			t.Errorf("Expected: %d, Got: %d", expectedTotal, total)
		}
	})
}

func TestCountWords(t *testing.T) {
	t.Run("returns the total words in a reader", func(t *testing.T) {
		expectedTotal := 15
		total, _ := countWords(strings.NewReader(`Hello, this is my test file.
To Test the line count.
It should be 3`))

		if total != expectedTotal {
			t.Errorf("Expected: %d, Got: %d", expectedTotal, total)
		}
	})
}

func TestCountCharacters(t *testing.T) {
	t.Run("returns the total characters in a reader", func(t *testing.T) {
		expectedTotal := 67
		total, _ := countCharacters(strings.NewReader(`Hello, this is my test file.
To Test the line count.
It should be 3`))

		if total != expectedTotal {
			t.Errorf("Expected: %d, Got: %d", expectedTotal, total)
		}
	})
}
