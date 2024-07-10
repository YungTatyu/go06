package piscine

import "os"

const (
	cmd      = "cat: "
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

func ReadPrintLoop() int {
	for {
		var buffer [buffsize]byte
		n, err := os.Stdin.Read(buffer[:])
		if n == 0 {
			return 0
		}
		if err != nil {
			PrintError(err)
			return 1
		}
		n, err = os.Stdout.Write(buffer[:])
		if err != nil {
			PrintError(err)
			return 1
		}
	}
}

func Cat() int {
	if argc == 1 {
		return ReadPrintLoop()
	}
	argv := os.Args[1:]
	var re int = 0
	for _, file := range argv {
		re |= DisplayFile(file)
	}
	return re
}
