package internals

import (
	"encoding/xml"
	"os"
)

type Filename struct {
    Module string `xml:"module,attr"`
    Name   string `xml:"files>filename"`
}

func findMainFileDirectory(path string, moduleName string) (string, error) {
	var xmlFileName string = path + "/" + moduleName + ".xml"
	var file Filename

	byteValue, err := os.ReadFile(xmlFileName)
	if err != nil {
		return "", err
	}

	xml.Unmarshal(byteValue, &file)
	return file.Name, nil
}