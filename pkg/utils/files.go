package utils

import "os"

func AppendToFile(src string, dst string) error {
	file, err := os.OpenFile(dst, os.O_APPEND, 6)
	if err != nil {
		return err
	}

	file.WriteString(src)
	return nil
}

func ReadFile(fileName string) (string, error) {
	b, err := os.ReadFile(fileName)
	if err != nil {
		return "", err
	}

	return string(b), nil
}