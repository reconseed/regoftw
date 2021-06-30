package utils

import (
	"log"
	"os"
)

type File interface {
	Close()
	WriteLn(data string)
	WriteLnFromList(data []string)
}

type file struct {
	f *os.File
}

var f File = nil

func GetFile(path string) File {
	return newFile(path)
}

func newFile(fileName string) File {
	var err error
	var openFile *os.File
	openFile, err = os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Failed creating file")
	}
	return &file{openFile}
}

func (f *file) WriteLnFromList(data []string) {
	var err error
	if f == nil {
		return
	}
	for _, d := range data {
		_, err = f.f.WriteString(d + "\n")
		if err != nil {
			continue
		}
	}
}

func (f *file) WriteLn(data string) {
	if f != nil {
		f.f.WriteString(data + "\n")
	}
}

func (f *file) Close() {
	if f != nil {
		f.f.Close()
	}
}
