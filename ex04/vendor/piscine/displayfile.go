package piscine

import (
	"os"
)

func DisplayFile(fileName string, bytes int) int {
	file, err := os.Open(fileName)
	if err != nil {
		PrintError(err)
		return 1
	}
	defer file.Close()
	filesize := calcFileSize(file)
	_, err = file.Seek(-int64(min(bytes, int(filesize))), os.SEEK_END)
	if err != nil {
		PrintError(err)
		return 1
	}
	buffer := make([]byte, min(bytes, int(filesize)))
	_, err = file.Read(buffer)
	if err != nil {
		PrintError(err)
		return 1
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
