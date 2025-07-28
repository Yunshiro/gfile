package gfile

import (
	"os"
	"testing"
)

func TestToFile(t *testing.T) {
	testData := []byte("Hello, Test!")
	tmpFile := "test_output.txt"

	// delete file after testing
	defer os.Remove(tmpFile)

	res := toFile(testData, tmpFile)
	if !res {
		t.Fatal("toFile() return false")
	}
	
	bts := Transfer("test_output.txt")
	if string(bts) != string(testData) {
		t.Errorf("file content not match.\nhope: %s\nfact: %s",  testData, bts)
	}
}