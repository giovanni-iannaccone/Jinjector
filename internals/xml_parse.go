package internals

import (
	"encoding/xml"
	"errors"
	"os"
)

type File struct {
	Module string `xml:"module,attr"`
	Dir    string `xml:",chardata"`
}

type Files struct {
	Filenames []File `xml:"filename"`
}

type Extension struct {
	Files Files `xml:"files"`
}

func findMainFileDirectory(path string, moduleName string) (string, error) {
	xmlFileName := path + "/" + moduleName + ".xml"
	var extension Extension

	byteValue, err := os.ReadFile(xmlFileName)
	if err != nil {
		return "", err
	}

	err = xml.Unmarshal(byteValue, &extension)
	if err != nil {
		return "", err
	}

	for _, file := range extension.Files.Filenames {
		if file.Module == moduleName {
			return file.Dir, nil
		}
	}

	return "", errors.New("can't find main file")
}
