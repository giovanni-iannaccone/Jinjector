package utils

import (
	"testing"
)

func TestFileWriteAndReading(t *testing.T) {
	var testString string = "test"

	AppendToFile(testString, "../../test/test.txt")
	buffer, _ := ReadFile("../../test/test.txt")

	if buffer != testString {
		t.Fatalf("Read and write => Failed: %s", buffer)
	}
}
