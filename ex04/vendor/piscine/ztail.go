package piscine

import "os"

const (
	cmd      = "tail: "
	buffsize = 1024
)

var argc = func() int {
	var argv []string = os.Args
	var length int = 0
	for range argv {
		length++
	}
	return length
}()

func PrintError(err error) {
	os.Stderr.WriteString(cmd + err.Error() + "\n")
}

func PrintErrorMsg(msg string) {
	os.Stderr.WriteString(cmd + msg + "\n")
}

func Usage(msg string) {
	os.Stderr.WriteString(cmd + msg + "\n")
}

func ZTail() int {
	if argc == 1 {
		Usage("requires two or more arguments")
		return 1
	}
	argv := os.Args[1:]
	var t Tail = Tail{}
	if !t.Parse(argv) {
		return 1
	}
	if t.bytes == 0 {
		return 0
	}
	var re int
	for i, file := range t.files {
		re |= DisplayFile(file, uint(i), &t)
	}
	return re
}
