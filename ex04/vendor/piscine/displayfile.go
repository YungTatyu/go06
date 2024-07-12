package piscine

import (
	"os"
)

func DisplayFile(fileName string, index uint, t *Tail) int {
	file, err := os.Open(fileName)
	if err != nil {
		PrintError(err)
		return 1
	}
	defer file.Close()
	filesize := calcFileSize(file)
	_, err = file.Seek(-int64(min(t.bytes, int(filesize))), os.SEEK_END)
	if err != nil {
		PrintError(err)
		return 1
	}
	buffer := make([]byte, min(t.bytes, int(filesize)))
	_, err = file.Read(buffer)
	if err != nil {
		PrintError(err)
		return 1
	}
	len := length(t.files)
	if len > 1 {
		buffer = append(createPrefix(fileName, index != 0), buffer...)
	}
	_, err = os.Stdout.Write(buffer[:])
	if err != nil {
		PrintError(err)
		return 1
	}
	return 0
}

func calcFileSize(f *os.File) int64 {
	fileinfo, err := f.Stat()
	if err != nil {
		PrintError(err)
		return -1
	}
	return fileinfo.Size()
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func createPrefix(fileName string, nl bool) []byte {
	prefix := "==> " + fileName + " <==\n"
	if nl {
		prefix = "\n" + prefix
	}
	return []byte(prefix)
}

func length(arr []string) uint {
	var i uint = 0
	for range arr {
		i++
	}
	return i
}
