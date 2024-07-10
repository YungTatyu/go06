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
		Usage("requires two or more arguments")
		return 1
	}
	argv := os.Args[1:]
	var t Tail = Tail{}
	if !t.Parse(argv) {
		return 1
	}

	// for _, file := range argv {
	// 	re |= DisplayFile(file)
	// }
	return 0
}
