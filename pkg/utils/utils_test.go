package utils

import (
	"os"
	"testing"
)

func TestFileWriteAndReading(t *testing.T) {
	var testString string = "test"
	var testFile string = "../../test/test.txt"
	AppendToFile(testString, testFile)
	buffer, _ := ReadFile(testFile)

	os.WriteFile(testFile, []byte(""), 2)
	if buffer != testString {
		t.Fatalf("Read and write => Failed: %s", buffer)
	}
}
