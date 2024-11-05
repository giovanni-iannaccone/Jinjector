package internals

import (
	"testing"
	"data"
)

func TestAddAddressToString(t *testing.T) {
	var addr data.Address = data.Address{Ip: "10.10.10.10", Port: 5555}
	var testFileContent string = "test"
	addAddressToString(addr, &testFileContent)

	if testFileContent != "$IP='10.10.10.10';$PORT=5555;test" {
		t.Fatalf("Adding address => Failed %s", testFileContent)
	}
}

func TestFindModuleName(t *testing.T) {
	var path string = "test/module/name/fake_module_name"
	var moduleName = findModuleName(path)
	if moduleName != "fake_module_name" {
		t.Fatalf("Finding module's name => Failed: %s", moduleName)
	}
}

func TestFindMainFileDirectory(t *testing.T) {
	mainDirectory, _ := findMainFileDirectory("../test", "test")
	if mainDirectory != "test/directory/main.php" {
		t.Fatalf("Finding module's main file => Failed: %s", mainDirectory)
	}
}