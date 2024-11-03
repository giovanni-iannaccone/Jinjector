package utils

import (
	"data"
	"fmt"
)

func Print(color string, format string, args ...any) {
	fmt.Print(color)
	fmt.Printf(format, args...)
	fmt.Print(data.Reset)
}