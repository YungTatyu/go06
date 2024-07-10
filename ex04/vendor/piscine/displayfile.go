package piscine

import (
	"os"
)

func DisplayFile(file string) int {
	data, err := os.ReadFile(file)
	if err != nil {
		PrintError(err)
		return 1
	}
	_, err = os.Stdout.Write(data)
	if err != nil {
		PrintError(err)
		return 1
	}
	return 0
}
