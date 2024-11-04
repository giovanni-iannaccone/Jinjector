package internals

import "testing"

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