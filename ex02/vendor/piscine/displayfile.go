package piscine

import (
	"os"
)

var argc = func() int {
	var argv []string = os.Args
	var length int = 0
	for range argv {
		length++
	}
	return length
}()

func DisplayFile() {
	if argc == 1 {
		os.Stderr.WriteString("File name missing\n")
		return
	}
	if argc != 2 {
		os.Stderr.WriteString("Too many arguments\n")
		return
	}
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		os.Stderr.WriteString("failed to open a file\n")
		return
	}
	_, err = os.Stdout.Write(data)
	if err != nil {
		os.Stderr.WriteString("failed to write to stdout: " + err.Error() + "\n")
	}
}
